import { useEffect, useState } from "react";

import { useNavigate, useParams } from "react-router-dom";
import RegisterSuccessRender from "./RegisterSuccessPage";


export default function Register() {
  const [refCodeValid, setRefCodeValid] = useState(false);
  const [checking, setChecking] = useState(true);

  const params = useParams();
  const refCode = params.id

  const navigate = useNavigate()

  

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
  if (!checking) {
    if (!refCodeValid) {
      navigate(`/register?code=-1`)
      return
    } 
    return <RegisterSuccessRender />
  }
}

async function checkRefCode(referral: string): Promise<boolean> {
  const verifyApiPath = "/api/verify/referral-code";
  const response = await fetch(verifyApiPath, {
    method: "GET",
    headers: {
      referralCode: referral.trim(),
    },
  });
  return response.ok;
}
