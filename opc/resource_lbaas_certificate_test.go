package opc

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccLBaaSServerCertificate_ServerCertificate(t *testing.T) {
	if checkSkipLBTests() {
		t.Skip(fmt.Printf("`OPC_LBAAS_ENDPOINT` not set, skipping test"))
	}

	rInt := acctest.RandInt()
	resName := "opc_lbaas_certificate.server-cert"
	testName := fmt.Sprintf("acctest-%d", rInt)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: opcResourceCheck(resName, testAccLBaaSCheckCertificateDestroyed),
		Steps: []resource.TestStep{
			{
				Config: testAccLBaaSCertificateConfig_ServerCertificate(rInt),
				Check: resource.ComposeTestCheckFunc(
					opcResourceCheck(resName, testAccLBaaSCheckCertificateExists),
					resource.TestCheckResourceAttr(resName, "name", testName),
					resource.TestCheckResourceAttr(resName, "type", "SERVER"),
					resource.TestMatchResourceAttr(resName, "uri", regexp.MustCompile(testName)),
					resource.TestCheckResourceAttrSet(resName, "certificate_body"),
					resource.TestCheckResourceAttrSet(resName, "certificate_chain"),
				),
			},
		},
	})
}

func TestAccLBaaSServerCertificate_TrustedCertificate(t *testing.T) {
	if checkSkipLBTests() {
		t.Skip(fmt.Printf("`OPC_LBAAS_ENDPOINT` not set, skipping test"))
	}

	rInt := acctest.RandInt()
	resName := "opc_lbaas_certificate.trusted-cert"
	testName := fmt.Sprintf("acctest-%d", rInt)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: opcResourceCheck(resName, testAccLBaaSCheckCertificateDestroyed),
		Steps: []resource.TestStep{
			{
				Config: testAccLBaaSCertificateConfig_TrustedCertificate(rInt),
				Check: resource.ComposeTestCheckFunc(
					opcResourceCheck(resName, testAccLBaaSCheckCertificateExists),
					resource.TestCheckResourceAttr(resName, "name", testName),
					resource.TestCheckResourceAttr(resName, "type", "TRUSTED"),
					resource.TestMatchResourceAttr(resName, "uri", regexp.MustCompile(testName)),
					resource.TestCheckResourceAttrSet(resName, "certificate_body"),
					resource.TestCheckResourceAttrSet(resName, "certificate_chain"),
				),
			},
		},
	})
}

func testAccLBaaSCertificateConfig_ServerCertificate(rInt int) string {
	return fmt.Sprintf(`
resource "opc_lbaas_certificate" "server-cert" {
  name = "acctest-%d"
  type = "SERVER"
  private_key = "-----BEGIN RSA PRIVATE KEY-----\n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\n-----END RSA PRIVATE KEY-----\n"
  certificate_body = "-----BEGIN CERTIFICATE-----\nMIIFjzCCA3egAwIBAgIRAPHidcMOfXqzFXEVlGi/148wDQYJKoZIhvcNAQELBQAw\nPDEVMBMGA1UEChMMc2Nyb3Nzb3JhY2xlMSMwIQYDVQQDExpteXdlYmFwcC5zY3Jv\nc3NvcmFjbGUuc2l0ZTAeFw0xODA2MzAxMTM5MTVaFw0xOTA3MDQxNTM5MTVaMFsx\nCzAJBgNVBAYTAkNBMRAwDgYDVQQIEwdPbnRhcmlvMRUwEwYDVQQKEwxzY3Jvc3Nv\ncmFjbGUxIzAhBgNVBAMTGm15d2ViYXBwLnNjcm9zc29yYWNsZS5zaXRlMIICIjAN\nBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAtl8cFy2b92+hop3p2DU1gp3OnQf6\nc7njsrRYfKRaCHd7L93aiVKGkccFNNHVzhBc/bScfEhYtqFCstMPBNVgTqqzulF0\nifoNDRxg2Dwbk8SsO0aWzmz38vEew9xLdLE4KAze0DVfgPB/nrr75osBCoLnxC1T\nRBMMzp26Ff/0S1W9k/D8zUGezOPKH26x8hUdQaVxS+wZQQrWfvsWkBgFdDvysZXi\nPdpo8v09neeh616FNrrT6Z9LP2VDvc3YToqFaDgV6Yy0Qm7zwz4IOpc6NQ1sP/Mx\ns9f36Hb+k1KXvAFt/dB5ROeCYDvQIMbEN+tjTENMRukjyfgc5MOb8qViSANPl6pe\nZtloP92l3bYOD8LzXFfqAGtaY0XBEn1Aw0TLWmlVPjyrsv7uXSIlCm53YOH0FtQX\nq6oikQCj0/S/Z+D0LpOM+JElvRb4JjiZIqGSVpIgxcZs23xt3GyQMAEm64CF5/y6\nIsB/E9W1Ue10u1/41CcPpc9pzS8QbxwJyIQPuZHljl/R7RpZMf9m8VU2/lwYfZfI\n5iGbdTA9oc/t62KFHtk/uxh9fG1V9T30AtlpvP6sHKdSW9AKjPlcrV+qxvdQQznW\nNI8HuaBO+ynducTnfCCVBC05okFQge2NlFPE0p47w8wu4s2Yv/OICxiXZsm6gERu\nDJqxBFWnOvj2HFECAwEAAaNtMGswEwYDVR0lBAwwCgYIKwYBBQUHAwEwDAYDVR0T\nAQH/BAIwADAfBgNVHSMEGDAWgBSJNATOWcuXm+Jv0H7UCkIId2M0WTAlBgNVHREE\nHjAcghpteXdlYmFwcC5zY3Jvc3NvcmFjbGUuc2l0ZTANBgkqhkiG9w0BAQsFAAOC\nAgEAqcTBLaW4D5PcEYSwMhNYqdACCuV6mc1o18PzIDHn/VDqF8pVzvaSTdEuMTte\noz5W8JwBpG6jH8E3YKEMMC4f/CI09PdDM8nBr4yDOHlaTIt1jRWFjG7gBGfe6rZw\ntBkrz9fteU80ST8LBcEFnoov7Txss54amS0L+vXU1ddwx6e6k9Ta3eAWMNn/JkVg\n63uCgiYueLT62AJUJZvBwPJdBnYASpJxh/AN8biIkWqWnoERVoofGfngiGoQ9DLo\niq6So6Ix4D95eOmIRpf/MC2yTTOeIxiQXi5LMk/NZ9oRUbc0JOinMVFfLStynnlP\n6xu0RjBKtCO2EjiRWl8sdQIVEgY3MicDoCoBt5HwJdIBkR955l4or5aFjxuLeV2E\nn2q2RacaUqV/xp46RCjm1hbJYkwcWXHnQzHkx6Jk6Y5kDcYp85CkUnBfzdfsC6/I\nbKByd3Sfp2wWvB9f+D1rI2ZsfOc18N/S+9AM/SVn0WHIr1DZBm2yaABe1m5Qo4jq\nXlRYAfjVg+tdhEwy6X6V8APW2ZLKXoVHTBl6XEc0Kqgths3r551nIHci9S1skMaN\ngO09dPxaJG+TcDKkO3YaxppszSY/IJa+h4nFv6j/mr+3tYRJ+Qs70tu+pqm2qcXI\noqW2x+i+3LFQD4vclv0aqrhVfeerpiSXpcGGyg65p9KtBCg=\n-----END CERTIFICATE-----\n"
  certificate_chain = "-----BEGIN CERTIFICATE-----\nMIIFazCCA1OgAwIBAgIQIPffd3HSy1azL08m3+003TANBgkqhkiG9w0BAQsFADA8\nMRUwEwYDVQQKEwxzY3Jvc3NvcmFjbGUxIzAhBgNVBAMTGm15d2ViYXBwLnNjcm9z\nc29yYWNsZS5zaXRlMB4XDTE4MDYzMDExMzkxNVoXDTE5MDcwNDE1MzkxNVowPDEV\nMBMGA1UEChMMc2Nyb3Nzb3JhY2xlMSMwIQYDVQQDExpteXdlYmFwcC5zY3Jvc3Nv\ncmFjbGUuc2l0ZTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBALZfHBct\nm/dvoaKd6dg1NYKdzp0H+nO547K0WHykWgh3ey/d2olShpHHBTTR1c4QXP20nHxI\nWLahQrLTDwTVYE6qs7pRdIn6DQ0cYNg8G5PErDtGls5s9/LxHsPcS3SxOCgM3tA1\nX4Dwf566++aLAQqC58QtU0QTDM6duhX/9EtVvZPw/M1Bnszjyh9usfIVHUGlcUvs\nGUEK1n77FpAYBXQ78rGV4j3aaPL9PZ3noetehTa60+mfSz9lQ73N2E6KhWg4FemM\ntEJu88M+CDqXOjUNbD/zMbPX9+h2/pNSl7wBbf3QeUTngmA70CDGxDfrY0xDTEbp\nI8n4HOTDm/KlYkgDT5eqXmbZaD/dpd22Dg/C81xX6gBrWmNFwRJ9QMNEy1ppVT48\nq7L+7l0iJQpud2Dh9BbUF6uqIpEAo9P0v2fg9C6TjPiRJb0W+CY4mSKhklaSIMXG\nbNt8bdxskDABJuuAhef8uiLAfxPVtVHtdLtf+NQnD6XPac0vEG8cCciED7mR5Y5f\n0e0aWTH/ZvFVNv5cGH2XyOYhm3UwPaHP7etihR7ZP7sYfXxtVfU99ALZabz+rByn\nUlvQCoz5XK1fqsb3UEM51jSPB7mgTvsp3bnE53wglQQtOaJBUIHtjZRTxNKeO8PM\nLuLNmL/ziAsYl2bJuoBEbgyasQRVpzr49hxRAgMBAAGjaTBnMA4GA1UdDwEB/wQE\nAwICBDAPBgNVHRMBAf8EBTADAQH/MB0GA1UdDgQWBBSJNATOWcuXm+Jv0H7UCkII\nd2M0WTAlBgNVHREEHjAcghpteXdlYmFwcC5zY3Jvc3NvcmFjbGUuc2l0ZTANBgkq\nhkiG9w0BAQsFAAOCAgEAVREHaF6qPPe27c/IMThVLDDzihaT83DWjzcel/OI4MPP\nsd1HV01tWtt56jP4IjzH1SpdHQMncTbSAbEOwsbclmrLe4E/Hd58Dzjo6apkKx2M\nieX+XVBi0KJ5pKh+OJHug8CGpnFu7IWla5zUiRY2Mm4Y3EdZNn4NH0smd8Expqck\nqehI0xsuN4blj3KFRtmgN1Zm48qSavah9PfGpicCPs1ZvoJJ8v17DFE4uFbkGqZl\nRRFpCybOmW7KeU5v8lDhmkcP6bu72xw+J7VGT0TfHotXTLXSPNjRlD13m1idvk0o\nXiosdLoQWvMpq71mstG3b11fwCA/EXuJkgANTxTkpjo7S5fWvgDUPqaVTt9nSbt0\nPHID1OvxfKZeuNIB0hM0oA+C5ZbSULuWTEaHPIwM3xgM+I7gCoJItJpzruyrtSjE\nUNJlMlo9zoJptx/a6ZguIvyu95MQbDnTJfq8sZjK1r0mxMBvx9tE8qTHXgAkuIC3\nFpDuFtfIDUgiWweSk5js19/deiP+tQ2abd/Z8MCR++e0bHNMdyyXS9CahOcSWCCJ\nHomAUmji594MTlP37MfkufA9NGegIwACf0VqE6FWrriO6VThvNpjnkNBewttlymu\nRshjxzWs/8bVI3HyOIz3CVh2gD2477D+kDsJJICBxkmz2eizt8EUcZwVsXO0KRs=\n-----END CERTIFICATE-----\n"
}
`, rInt)
}

func testAccLBaaSCertificateConfig_TrustedCertificate(rInt int) string {
	return fmt.Sprintf(`
resource "opc_lbaas_certificate" "trusted-cert" {
  name = "acctest-%d"
  type = "TRUSTED"
  certificate_body = "-----BEGIN CERTIFICATE-----\nMIIFjzCCA3egAwIBAgIRAPHidcMOfXqzFXEVlGi/148wDQYJKoZIhvcNAQELBQAw\nPDEVMBMGA1UEChMMc2Nyb3Nzb3JhY2xlMSMwIQYDVQQDExpteXdlYmFwcC5zY3Jv\nc3NvcmFjbGUuc2l0ZTAeFw0xODA2MzAxMTM5MTVaFw0xOTA3MDQxNTM5MTVaMFsx\nCzAJBgNVBAYTAkNBMRAwDgYDVQQIEwdPbnRhcmlvMRUwEwYDVQQKEwxzY3Jvc3Nv\ncmFjbGUxIzAhBgNVBAMTGm15d2ViYXBwLnNjcm9zc29yYWNsZS5zaXRlMIICIjAN\nBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAtl8cFy2b92+hop3p2DU1gp3OnQf6\nc7njsrRYfKRaCHd7L93aiVKGkccFNNHVzhBc/bScfEhYtqFCstMPBNVgTqqzulF0\nifoNDRxg2Dwbk8SsO0aWzmz38vEew9xLdLE4KAze0DVfgPB/nrr75osBCoLnxC1T\nRBMMzp26Ff/0S1W9k/D8zUGezOPKH26x8hUdQaVxS+wZQQrWfvsWkBgFdDvysZXi\nPdpo8v09neeh616FNrrT6Z9LP2VDvc3YToqFaDgV6Yy0Qm7zwz4IOpc6NQ1sP/Mx\ns9f36Hb+k1KXvAFt/dB5ROeCYDvQIMbEN+tjTENMRukjyfgc5MOb8qViSANPl6pe\nZtloP92l3bYOD8LzXFfqAGtaY0XBEn1Aw0TLWmlVPjyrsv7uXSIlCm53YOH0FtQX\nq6oikQCj0/S/Z+D0LpOM+JElvRb4JjiZIqGSVpIgxcZs23xt3GyQMAEm64CF5/y6\nIsB/E9W1Ue10u1/41CcPpc9pzS8QbxwJyIQPuZHljl/R7RpZMf9m8VU2/lwYfZfI\n5iGbdTA9oc/t62KFHtk/uxh9fG1V9T30AtlpvP6sHKdSW9AKjPlcrV+qxvdQQznW\nNI8HuaBO+ynducTnfCCVBC05okFQge2NlFPE0p47w8wu4s2Yv/OICxiXZsm6gERu\nDJqxBFWnOvj2HFECAwEAAaNtMGswEwYDVR0lBAwwCgYIKwYBBQUHAwEwDAYDVR0T\nAQH/BAIwADAfBgNVHSMEGDAWgBSJNATOWcuXm+Jv0H7UCkIId2M0WTAlBgNVHREE\nHjAcghpteXdlYmFwcC5zY3Jvc3NvcmFjbGUuc2l0ZTANBgkqhkiG9w0BAQsFAAOC\nAgEAqcTBLaW4D5PcEYSwMhNYqdACCuV6mc1o18PzIDHn/VDqF8pVzvaSTdEuMTte\noz5W8JwBpG6jH8E3YKEMMC4f/CI09PdDM8nBr4yDOHlaTIt1jRWFjG7gBGfe6rZw\ntBkrz9fteU80ST8LBcEFnoov7Txss54amS0L+vXU1ddwx6e6k9Ta3eAWMNn/JkVg\n63uCgiYueLT62AJUJZvBwPJdBnYASpJxh/AN8biIkWqWnoERVoofGfngiGoQ9DLo\niq6So6Ix4D95eOmIRpf/MC2yTTOeIxiQXi5LMk/NZ9oRUbc0JOinMVFfLStynnlP\n6xu0RjBKtCO2EjiRWl8sdQIVEgY3MicDoCoBt5HwJdIBkR955l4or5aFjxuLeV2E\nn2q2RacaUqV/xp46RCjm1hbJYkwcWXHnQzHkx6Jk6Y5kDcYp85CkUnBfzdfsC6/I\nbKByd3Sfp2wWvB9f+D1rI2ZsfOc18N/S+9AM/SVn0WHIr1DZBm2yaABe1m5Qo4jq\nXlRYAfjVg+tdhEwy6X6V8APW2ZLKXoVHTBl6XEc0Kqgths3r551nIHci9S1skMaN\ngO09dPxaJG+TcDKkO3YaxppszSY/IJa+h4nFv6j/mr+3tYRJ+Qs70tu+pqm2qcXI\noqW2x+i+3LFQD4vclv0aqrhVfeerpiSXpcGGyg65p9KtBCg=\n-----END CERTIFICATE-----\n"
  certificate_chain = "-----BEGIN CERTIFICATE-----\nMIIFazCCA1OgAwIBAgIQIPffd3HSy1azL08m3+003TANBgkqhkiG9w0BAQsFADA8\nMRUwEwYDVQQKEwxzY3Jvc3NvcmFjbGUxIzAhBgNVBAMTGm15d2ViYXBwLnNjcm9z\nc29yYWNsZS5zaXRlMB4XDTE4MDYzMDExMzkxNVoXDTE5MDcwNDE1MzkxNVowPDEV\nMBMGA1UEChMMc2Nyb3Nzb3JhY2xlMSMwIQYDVQQDExpteXdlYmFwcC5zY3Jvc3Nv\ncmFjbGUuc2l0ZTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBALZfHBct\nm/dvoaKd6dg1NYKdzp0H+nO547K0WHykWgh3ey/d2olShpHHBTTR1c4QXP20nHxI\nWLahQrLTDwTVYE6qs7pRdIn6DQ0cYNg8G5PErDtGls5s9/LxHsPcS3SxOCgM3tA1\nX4Dwf566++aLAQqC58QtU0QTDM6duhX/9EtVvZPw/M1Bnszjyh9usfIVHUGlcUvs\nGUEK1n77FpAYBXQ78rGV4j3aaPL9PZ3noetehTa60+mfSz9lQ73N2E6KhWg4FemM\ntEJu88M+CDqXOjUNbD/zMbPX9+h2/pNSl7wBbf3QeUTngmA70CDGxDfrY0xDTEbp\nI8n4HOTDm/KlYkgDT5eqXmbZaD/dpd22Dg/C81xX6gBrWmNFwRJ9QMNEy1ppVT48\nq7L+7l0iJQpud2Dh9BbUF6uqIpEAo9P0v2fg9C6TjPiRJb0W+CY4mSKhklaSIMXG\nbNt8bdxskDABJuuAhef8uiLAfxPVtVHtdLtf+NQnD6XPac0vEG8cCciED7mR5Y5f\n0e0aWTH/ZvFVNv5cGH2XyOYhm3UwPaHP7etihR7ZP7sYfXxtVfU99ALZabz+rByn\nUlvQCoz5XK1fqsb3UEM51jSPB7mgTvsp3bnE53wglQQtOaJBUIHtjZRTxNKeO8PM\nLuLNmL/ziAsYl2bJuoBEbgyasQRVpzr49hxRAgMBAAGjaTBnMA4GA1UdDwEB/wQE\nAwICBDAPBgNVHRMBAf8EBTADAQH/MB0GA1UdDgQWBBSJNATOWcuXm+Jv0H7UCkII\nd2M0WTAlBgNVHREEHjAcghpteXdlYmFwcC5zY3Jvc3NvcmFjbGUuc2l0ZTANBgkq\nhkiG9w0BAQsFAAOCAgEAVREHaF6qPPe27c/IMThVLDDzihaT83DWjzcel/OI4MPP\nsd1HV01tWtt56jP4IjzH1SpdHQMncTbSAbEOwsbclmrLe4E/Hd58Dzjo6apkKx2M\nieX+XVBi0KJ5pKh+OJHug8CGpnFu7IWla5zUiRY2Mm4Y3EdZNn4NH0smd8Expqck\nqehI0xsuN4blj3KFRtmgN1Zm48qSavah9PfGpicCPs1ZvoJJ8v17DFE4uFbkGqZl\nRRFpCybOmW7KeU5v8lDhmkcP6bu72xw+J7VGT0TfHotXTLXSPNjRlD13m1idvk0o\nXiosdLoQWvMpq71mstG3b11fwCA/EXuJkgANTxTkpjo7S5fWvgDUPqaVTt9nSbt0\nPHID1OvxfKZeuNIB0hM0oA+C5ZbSULuWTEaHPIwM3xgM+I7gCoJItJpzruyrtSjE\nUNJlMlo9zoJptx/a6ZguIvyu95MQbDnTJfq8sZjK1r0mxMBvx9tE8qTHXgAkuIC3\nFpDuFtfIDUgiWweSk5js19/deiP+tQ2abd/Z8MCR++e0bHNMdyyXS9CahOcSWCCJ\nHomAUmji594MTlP37MfkufA9NGegIwACf0VqE6FWrriO6VThvNpjnkNBewttlymu\nRshjxzWs/8bVI3HyOIz3CVh2gD2477D+kDsJJICBxkmz2eizt8EUcZwVsXO0KRs=\n-----END CERTIFICATE-----\n"
}
`, rInt)
}

func testAccLBaaSCheckCertificateExists(state *OPCResourceState) error {
	name := state.Attributes["name"]

	client := testAccProvider.Meta().(*Client).lbaasClient.SSLCertificateClient()

	if _, err := client.GetSSLCertificate(name); err != nil {
		return fmt.Errorf("Error retrieving state of Certificate '%s': %v", name, err)
	}

	return nil
}

func testAccLBaaSCheckCertificateDestroyed(state *OPCResourceState) error {
	name := state.Attributes["name"]

	client := testAccProvider.Meta().(*Client).lbaasClient.SSLCertificateClient()

	if info, _ := client.GetSSLCertificate(name); info != nil {
		return fmt.Errorf("Certificate '%s' still exists: %+v", name, info)
	}
	return nil
}
