interface IDoRegister {
  email: string;
  password: string; //hash before request
  securityText: string; //process into ecdhPrivate
  referralCode: string; //use to update used status of referral code
  ecdhPrivate: string;
}

export default function DoRegister(iDoRegister: IDoRegister) { //call after checking password repeat
  const jsonBody = {
    email: iDoRegister.email,
    referralCode: iDoRegister.referralCode,
  };




}
