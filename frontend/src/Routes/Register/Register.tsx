import { useEffect, useState } from "react";

import { useParams } from "react-router-dom";
import RegisterSuccess from "./RegisterSuccessPage";

export default function Register() {
  const [refCodeValid, setRefCodeValid] = useState(false);
  const [checking, setChecking] = useState(true);

  const { refCode } = useParams();

  useEffect(() => {
    if (typeof refCode !== "string") {
      setRefCodeValid(false);
      setChecking(false);
      return;
    }
    const verifyCode = async (code: string) => {
      const validity = await checkRefCode(code);
      if (!validity) {
        setRefCodeValid(false);
      } else {
        setRefCodeValid(true);
      }
      setChecking(false);
    };
    verifyCode(refCode);
  }, []);
  if (checking) {
    return <div>Loading...</div>
  }
  if (!checking) {
    return refCodeValid ? <RegisterSuccess /> : <div>Failed sad </div>
  }
}

async function checkRefCode(referral: string): Promise<boolean> {
  //TODO
  const verifyApiPath = "/api/user/protected/verify-referral";
  var response = await fetch(verifyApiPath, {
    method: "POST",
    headers: {
      referralCode: referral.trim(),
    },
  });
  return response.ok;
}
