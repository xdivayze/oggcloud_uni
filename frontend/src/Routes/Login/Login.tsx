import { useEffect, useRef, useState } from "react";
import Navbar from "../../Navbar/Navbar";
import ObeseBar from "../Register/Components/ObeseBar";

import ComponentDispatchStruct from "../Register/Components/ComponentDispatchStruct";
import {
  GetSaveUserText,
  ObeseBarDefaultStyles,
} from "../Register/Services/utils";
import { ValidatePassword } from "./Service/PasswordService";
import { DoCheckMailValidity } from "../Register/Services/MailServices";
import { SendLoginRequest } from "./Service/Login";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../../Protected/AuthProvider";
import SubmitButton from "../Register/Components/SubmitButton";
import { SAVED_FIELDNAME } from "./Service/constants";

export default function Login() {
  const auth = useAuth();

  const [save, setSave] = useState(false);
  
  useEffect(() => {}, []); //TODO check for saved sign-in

  const emailCompStruct = ComponentDispatchStruct(
    ObeseBarDefaultStyles,
    "Enter your email(e.g. example@example.org)"
  );

  const passwordCompStruct = ComponentDispatchStruct(
    ObeseBarDefaultStyles,
    "Enter your password"
  );

  const securityTextCompStruct = ComponentDispatchStruct(
    ObeseBarDefaultStyles,
    "Enter your security txt"
  );

  const navigate = useNavigate();

  const onSubmitClick = async () => {
    try {
      ValidatePassword(passwordCompStruct);
      if (!DoCheckMailValidity(emailCompStruct)) {
        throw new Error("mail invalid");
      }
    } catch (e: any) {
      throw e;
    }
    SendLoginRequest(
      passwordCompStruct.getRefContent().innerText,
      emailCompStruct.getRefContent().innerText
    )
      .catch((e: Error) => {
        navigate("/err?message=" + e.message.trim());
        throw e;
      })
      .then((a: string) => {
        auth.login(a);
        save
          ? window.localStorage.setItem(
              SAVED_FIELDNAME,
              GetSaveUserText(emailCompStruct.getRefContent().innerText)
            )
          : void 0;
        navigate("/user/profile"); //TODO put under protected layout
      });
  };

  return (
    <div className="w-full mx-7">
      <div className="pt-14 min-h-1/5">
        <Navbar collapse={true} />
      </div>
      <div className="w-full">
        <p className="text-5xl font-robotoSlab text-indigo-ogg flex justify-center pt-4">
          LOGIN
        </p>
      </div>
      <div className="pt-25">
        <div className="flex flex-col ">
          <div className="px-40 flex flex-row w-full space-x-[300px]">
            <div className="w-1/2">
              <ObeseBar
                height="min-h-[110px]"
                color={emailCompStruct.styles}
                refPassed={emailCompStruct.getRef()}
                text={emailCompStruct.text}
                contentEditable
              />
            </div>
            <div className="w-1/2">
              <ObeseBar
                text={passwordCompStruct.text}
                color={passwordCompStruct.styles}
                refPassed={passwordCompStruct.getRef()}
                height="min-h-[110px]"
                contentEditable
              />
            </div>
          </div>
          <div className="px-40 flex flex-row w-full space-x-[300px]">
            <div className="w-1/2 mt-6">
              <ObeseBar
                refPassed={securityTextCompStruct.getRef()}
                height="min-h-[370px]"
                color={securityTextCompStruct.styles}
                text={securityTextCompStruct.text}
                contentEditable={true}
              />
            </div>
            <div className="w-1/2 mt-auto mb-2">
              <SubmitButton onSubmitClick={onSubmitClick} setSave={setSave} />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
