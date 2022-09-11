package main

import (
	"algatux/wallet/internal/authenticator"
	"algatux/wallet/internal/wallet"
)

const listenTo string = ":8080"
const rsaPrivateKey = "-----BEGIN RSA PRIVATE KEY-----\nMIICXAIBAAKBgFRriTZn2Xs4UDKeEKCZauNh3vp5usvpRqCnpIo5cbU1FYcOrODz\nLIwISwm/kB2LDCafOfDvpG3yvEIGtXSv1kkmTFti87VZK4iL3rrfiXrq+g2jylrs\nPhTUzAn2O/X70rquiGIOuyNfJRwp9ayInKKrAgBGCr9rYaAIA2zHGAHDAgMBAAEC\ngYBMyLa35oM4gULolz52ZLSE9usSFZBikd4sl+6fzpnvAMaA3kc+H9Bf2dcuma9i\nP6ugoWjZDY8YdhTnVSTManbTafS16rAeJrLp3mrAm33oxyp7xtoYEu1GMMyqi9rn\nZN4w/jxKQ6/gLmBxDorEH3/RIrSvD2ebuc1KFK1TSKE9eQJBAKiOAfsuUmgx3LXb\nCiOFNSL0LP467ir1M1atfirAtwaOW67jZ/PiDUHGvr042Ee0JNv9JhrOOeBX1xAp\nBTQUHvcCQQCAN3wavNly4kwbU9QYdyl4fTjOUSt7AWeHSf1C+4ayaxRTXPqwZVQw\nRWB09DxihufVNDtaK3ioSxgmGtVJIuSVAkEAp5XJ0nuT45Tv5MALrJVc54vu0Da3\nZm60xJFqyAcj8pjH/3KKgKlYlPWN34UNRP2PSErCABTa5ntvkNm4GreGsQJAYTyz\nC/wI1U219kue4GcOtmgROrboSMMJ5tpADhp/TrRSl9496KUQMgOLYWcQnJA/JbwU\n/w9U1B0PVWaoNm0V+QJBAI1Ls+JEFI5RopjaAUhNpgQav8bzDt9nf3Go9piL7Jfj\ne90hGAIqJuNlBZgGNPuFrWfnXlZrFOi9aewewYO75Y0=\n-----END RSA PRIVATE KEY-----"
const rsaPublicKey = "-----BEGIN PUBLIC KEY-----\nMIGeMA0GCSqGSIb3DQEBAQUAA4GMADCBiAKBgFRriTZn2Xs4UDKeEKCZauNh3vp5\nusvpRqCnpIo5cbU1FYcOrODzLIwISwm/kB2LDCafOfDvpG3yvEIGtXSv1kkmTFti\n87VZK4iL3rrfiXrq+g2jylrsPhTUzAn2O/X70rquiGIOuyNfJRwp9ayInKKrAgBG\nCr9rYaAIA2zHGAHDAgMBAAE=\n-----END PUBLIC KEY-----"

func main() {
	w := wallet.Wallet{}
	w.Boot(
		authenticator.RsaKeyPair{
			Private: rsaPrivateKey,
			Public:  rsaPublicKey,
		},
	).Start(listenTo)
}
