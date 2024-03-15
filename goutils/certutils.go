package goutils

import (
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"time"
)

// 常量定义
const (
	caCertFile     = "ca.crt"          // ca证书名称
	certKeyFile    = "cert.key"        // 服务端证书文件名
	certDecKeyFile = "cert_dec.key"    // 生死key文件名，文件由代码生成
	certPemFile    = "cert.pem"        // pem格式的证书私钥文件名
	rsaType        = "RSA PRIVATE KEY" // rsa
	privateKeyPerm = 0600              // 默认perm
)

// GetPrivateKey 获取解密后的密码,调用点负责使用完毕后清零内存
func GetPrivateKey(keyFile string) ([]byte, error) {
	fileName := filepath.Clean(keyFile)
	priKey, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return priKey, nil
}

// DecodePrivateKeyAndWriteNewFile 解密私钥文件
func DecodePrivateKeyAndWriteNewFile(paas, key []byte, newFile string) error {
	keyBlock, _ := pem.Decode(key)
	if keyBlock == nil {
		return errors.New("decode private key to block failed")
	}

	keyData, err := x509.DecryptPEMBlock(keyBlock, paas)
	if err != nil {
		return errors.New("DecryptPEMBlock failed")
	}

	plainKeyBlock := &pem.Block{
		Type:  rsaType,
		Bytes: keyData,
	}
	keyContent := pem.EncodeToMemory(plainKeyBlock)
	if keyContent == nil {
		return errors.New("private key encode to memory failed")
	}

	fileName := filepath.Clean(newFile)
	err = ioutil.WriteFile(fileName, keyContent, privateKeyPerm)
	if err != nil {
		return errors.New("write new keyfile" + fileName + " failed.")
	}
	return nil
}

// ReadCert read cert
func ReadCert(rootPath, fileName string) (string, error) {
	var realFile string
	var err error
	if len(rootPath) != 0 && len(fileName) != 0 {
		realFile = filepath.Join(rootPath, fileName)
	} else {
		realFile, err = filepath.Abs(rootPath)
		if err != nil {
			return "", fmt.Errorf("failed to convert filePath: %s to absolute path", rootPath)
		}
	}

	content, err := ioutil.ReadFile(realFile)
	if err != nil {
		// log.Error("read certificate err: %v", err)
		return "", err
	}
	return string(content), nil
}

// WriteCert write cert
func WriteCert(rootPath, fileName, certCxt string) error {
	var realFile string
	var err error
	if len(rootPath) != 0 && len(fileName) != 0 {
		realFile = filepath.Join(rootPath, fileName)
	} else {
		realFile, err = filepath.Abs(rootPath)
		if err != nil {
			return fmt.Errorf("failed to convert filePath: %s to absolute path", rootPath)
		}
	}

	err = ioutil.WriteFile(realFile, StringToByteSlice(certCxt), privateKeyPerm)
	if err != nil {
		// log.Error("write certificate err: %v", err)
	}
	return err
}

// WritePrivateKey write private key
func WritePrivateKey(filename, priKey, priPwd string) error {
	keyContent := DecryptPrivateKeyContent(priKey, priPwd)
	defer ClearByteArray(keyContent)
	if keyContent == nil {
		// log.Error("dec private key content failed")
		return errors.New("dec private key content failed")
	}

	if err := ioutil.WriteFile(filename, keyContent, privateKeyPerm); err != nil {
		// log.Error("write private key file err: %v", err)
		return err
	}
	return nil
}

// DecryptPrivateKeyContent 解密证书信息，得到证书私钥内容
func DecryptPrivateKeyContent(priKey, priPwd string) []byte {
	if priKey == "" || priPwd == "" {
		// log.Error("priKey or priPwd is nil")
		return nil
	}

	keyBlock, _ := pem.Decode(StringToByteSlice(priKey))
	if keyBlock == nil {
		// log.Error("decode key failed, not pem format")
		return nil
	}

	keyData, err := x509.DecryptPEMBlock(keyBlock, StringToByteSlice(priPwd))
	if err != nil {
		// log.Error("decode the encrypted key err: %v", err)
		return nil
	}
	defer ClearByteArray(keyData)

	plainKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: keyData,
	}
	return pem.EncodeToMemory(plainKeyBlock)
}

// IsValidCertAndKey check cert and key
func IsValidCertAndKey(cert, priKey, priPwd string) bool {
	if cert == "" || priKey == "" || priPwd == "" {
		// log.Error("has no cert, key or pwd")
		return false
	}
	return IsValidCert(cert)
}

// VerifyCACert return valid ca cert string
func VerifyCACert(certStr [][]string) string {
	validCertStr := ""
	for _, regResult := range certStr {
		if len(regResult) >= 1 && IsValidCert(regResult[0]) {
			if validCertStr != "" {
				validCertStr += "\n"
			}
			validCertStr += regResult[0]
		}
	}
	return validCertStr
}

// IsValidCert check cert and key
func IsValidCert(cert string) bool {
	if cert == "" {
		// log.Error("has no cert")
		return false
	}
	certBlock, _ := pem.Decode([]byte(cert))
	if certBlock == nil {
		// log.Error("decode cert failed, not pem format")
		return false
	}
	certInfo, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		// log.Error("parse cert err: %s", err.Error())
		return false
	}
	timeNow := time.Now()
	if timeNow.Before(certInfo.NotBefore) || timeNow.After(certInfo.NotAfter) {
		// log.Error("cert is not within the validity period")
		return false
	}
	return true
}

// ParseCAListByRex return valid ca by regex from input string
func ParseCAListByRex(caContent string) []*x509.Certificate {
	caRex := regexp.MustCompile(`-----BEGIN CERTIFICATE-----[\s\S]+?-----END CERTIFICATE-----`)
	caResults := caRex.FindAllStringSubmatch(caContent, -1)
	caList := make([]*x509.Certificate, 0)
	for _, caStr := range caResults {
		if len(caStr) >= 1 && IsValidCert(caStr[0]) {
			certBlock, _ := pem.Decode([]byte(caStr[0]))
			if certBlock == nil {
				continue
			}
			ca, err := x509.ParseCertificate(certBlock.Bytes)
			if err != nil {
				continue
			}
			caList = append(caList, ca)
		}
	}
	return caList
}

// GetCAListFromCaContentsByRex 从ca文件中解析出所有的ca块
func GetCAListFromCaContentsByRex(caContent string) []string {
	caRex := regexp.MustCompile(`-----BEGIN CERTIFICATE-----[\s\S]+?-----END CERTIFICATE-----`)
	caResults := caRex.FindAllStringSubmatch(caContent, -1)
	caList := make([]string, 0)
	for _, caStr := range caResults {
		if len(caStr) >= 1 && IsValidCert(caStr[0]) {
			caList = append(caList, caStr[0])
		}
	}
	return caList
}

// PrepareCerts for https client
func PrepareCerts(pemFile, newKeyFile, trustCaFile string) (*x509.CertPool, *tls.Certificate, error) {
	tlsByte, err := LoadCert(pemFile)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load tls cert, err %s", err.Error())
	}
	tlsKeyByte, err := LoadCert(newKeyFile)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load tls key, err: %s", err.Error())
	}

	defer ClearByteArray(tlsKeyByte)

	cert, err := tls.X509KeyPair(tlsByte, tlsKeyByte)
	if err != nil {
		return nil, nil, fmt.Errorf("generate certificate failed, err: %s", err.Error())
	}

	// ca file
	caBytes, err := LoadCert(trustCaFile)
	if err != nil {
		return nil, &cert, fmt.Errorf("failed to load ca file, err: %s", err.Error())
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caBytes)

	return certPool, &cert, err
}

// LoadCert Load cert content with given file path
func LoadCert(filePath string) ([]byte, error) {
	relPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to convert filePath: %s to absolute path", filePath)
	}

	cert, err := ioutil.ReadFile(relPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read cert file: %s, err: %s", relPath, err.Error())
	}

	return cert, nil
}

// GetValidCRL return vaild crl from input list
func GetValidCRL(crlList []string) map[string]*pkix.CertificateList {
	crlMap := make(map[string]*pkix.CertificateList)
	for _, crlContent := range crlList {
		crlTmp, err := x509.ParseCRL([]byte(crlContent))
		if err != nil {
			// log.Error("parse crl content err: %v, content: %s", err, crlContent)
			continue
		}
		if !crlTmp.HasExpired(time.Now()) {
			crlMap[crlContent] = crlTmp
		} else {
			// log.Error("crl has expired, content: %s", crlContent)
		}
	}
	return crlMap
}

// GetCAListByRex return vaild ca by redexp from input string
func GetCAListByRex(caContent string) []*x509.Certificate {
	caRex := regexp.MustCompile(`-----BEGIN CERTIFICATE-----[\s\S]+?-----END CERTIFICATE-----`)
	caResults := caRex.FindAllStringSubmatch(caContent, -1)
	caList := make([]*x509.Certificate, 0)
	for _, caStr := range caResults {
		if len(caStr) >= 1 && IsValidCert(caStr[0]) {
			certBlock, _ := pem.Decode([]byte(caStr[0]))
			if certBlock == nil {
				continue
			}
			ca, err := x509.ParseCertificate(certBlock.Bytes)
			if err != nil {
				continue
			}
			caList = append(caList, ca)
		}
	}
	return caList
}

// CheckCRLSignature check the signature of input crl by input ca cert
func CheckCRLSignature(caList []*x509.Certificate, crlMap map[string]*pkix.CertificateList) (string, int) {
	validCrl := ""
	validNum := 0
	for _, ca := range caList {
		for crlContent, crl := range crlMap {
			err := ca.CheckCRLSignature(crl)
			if err == nil {
				validCrl += crlContent
				validNum++
				delete(crlMap, crlContent)
			} else {
				// log.Error("crl not pass check signature, crl content: %s", crlContent)
			}
		}
	}
	return validCrl, validNum
}

func GenCertVerifyWithoutDNSName(pemFile, newKeyFile, trustCaFile string, rawCerts [][]byte) error {
	certPool, _, err := PrepareCerts(pemFile, newKeyFile, trustCaFile)
	if err != nil {
		// log.Error("initCertPool failed: " + err.Error())
		return err
	}

	if len(rawCerts) == 0 {
		// log.Error("no certificate to verify")
		return errors.New("no certificate to verify")
	}

	certs := make([]*x509.Certificate, len(rawCerts))
	for i, asn1Data := range rawCerts {
		cert, err := x509.ParseCertificate(asn1Data)
		if err != nil {
			// log.Error("x509 ParseCertificate fail, err: %s", err)
			return err
		}
		certs[i] = cert
	}

	opts := x509.VerifyOptions{
		Roots:         certPool,
		Intermediates: x509.NewCertPool(),
	}
	for _, cert := range certs[1:] {
		opts.Intermediates.AddCert(cert)
	}
	_, err = certs[0].Verify(opts)
	if err != nil {
		// log.Error("certs Verify fail, err: %s", err)
		return err
	}

	return nil
}
