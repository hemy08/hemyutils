package snowflake

import (
	"errors"
	"sync"
	"time"
)

const (
	workerBits  uint8 = 10 // 10bit 工作机器的id，如果你发现1024台机器不够那就调大次值
	numberBits  uint8 = 12 // 12bit 工作序号，如果你发现1毫秒并发生成4096个唯一id不够请调大次值
	workerMax   int64 = -1 ^ (-1 << workerBits)
	numberMax   int64 = -1 ^ (-1 << numberBits)
	timeShift   uint8 = workerBits + numberBits
	workerShift uint8 = numberBits
	// 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID，
	// 这个值请自行设置为你系统准备上线前的精确到毫秒级别的时间戳，因为雪花时间戳保证唯一的部分最多管69年（2的41次方），
	// 2021-11-15 19:53:50.224283 +0800 CST  fmt.Println(time.UnixMicro(t))
	startTime int64 = 1636977230224283
)

var (
	mu        sync.Mutex
	timestamp int64 = 0
	workerId  int64 = 1
	number    int64 = 0
)

func getSnowFlakeId() int64 {
	mu.Lock()
	defer mu.Unlock()
	now := time.Now().UnixNano() / 1e6
	if timestamp == now {
		number++
		if number > numberMax {
			for now <= timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		number = 0
		timestamp = now
	}
	// 以下表达式才是主菜
	//  (now-startTime)<<timeShift   产生了 41 + （10 + 12）的效应但却并不保证唯一
	//  | (w.workerId << workerShift)  保证了与其他机器不重复
	//  | (w.number))  保证了自己这台机不会重复
	ID := int64((now-startTime)<<timeShift | (workerId << workerShift) | (number))
	return ID
}

func GetID() int64 {
	return getSnowFlakeId()
}

// Worker 定义一个woker工作节点所需要的基本参数
type Worker struct {
	mu        sync.Mutex // 添加互斥锁 确保并发安全
	timestamp int64      // 记录上一次生成id的时间戳
	workerId  int64      // 该节点的ID
	number    int64      // 当前毫秒已经生成的id序列号(从0开始累加) 1毫秒内最多生成4096个ID
}

// NewWorker 实例化一个工作节点 workerId 为当前节点的id
func NewWorker(workerId int64) (*Worker, error) {
	// 要先检测workerId是否在上面定义的范围内
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("worker id excess of quantity")
	}
	// 生成一个新节点
	return &Worker{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

// GetID 生成方法一定要挂载在某个worker下，这样逻辑会比较清晰 指定某个节点生成id
func (w *Worker) GetID() int64 {
	// 获取id最关键的一点 加锁 加锁 加锁
	w.mu.Lock()
	defer w.mu.Unlock() // 生成完成后记得 解锁 解锁 解锁

	// 获取生成时的时间戳
	now := time.Now().UnixNano() / 1e6 // 纳秒转毫秒
	if w.timestamp == now {
		w.number++

		// 这里要判断，当前工作节点是否在1毫秒内已经生成numberMax个ID
		if w.number > numberMax {
			// 如果当前工作节点在1毫秒内生成的ID已经超过上限 需要等待1毫秒再继续生成
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		// 如果当前时间与工作节点上一次生成ID的时间不一致 则需要重置工作节点生成ID的序号
		w.number = 0
		// 下面这段代码看到很多前辈都写在if外面，无论节点上次生成id的时间戳与当前时间是否相同 都重新赋值  这样会增加一丢丢的额外开销 所以我这里是选择放在else里面
		w.timestamp = now // 将机器上一次生成ID的时间更新为当前时间
	}

	ID := int64((now-startTime)<<timeShift | (w.workerId << workerShift) | (w.number))
	return ID
}
