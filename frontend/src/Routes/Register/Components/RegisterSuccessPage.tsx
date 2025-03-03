import { useCallback, useRef, useState } from "react";
import ObeseBar from "./ObeseBar";
import Navbar from "../../../Navbar/Navbar";
import { IDoRegister } from "../Services/Register";
import { useParams } from "react-router-dom";
import { DoPasswordOperations } from "../Services/PasswordServices";
import { DoCheckMailValidity } from "../Services/MailServices";

export default function RegisterSuccess() {
  var emailRef = useRef<HTMLDivElement>(null);
  var passwordRef = useRef<HTMLDivElement>(null);
  var passwordRepeatRef = useRef<HTMLDivElement>(null);
  var securityTextRef = useRef<HTMLDivElement>(null);
  var submitRef = useRef<HTMLDivElement>(null);

  const [passwordText, setPasswordText] = useState(
    "Enter a password not over 9 characters"
  );
  const [passwordStyles, setPasswordStyles] = useState(
    "text-white bg-teal-ogg-1 hover:text-white hover:bg-indigo-950 text-2xl"
  );

  const [mailText, setMailText] = useState(
    "Enter your email(e.g. example@example.org)"
  );
  const [mailStyles, setMailStyles] = useState(
    "text-white bg-teal-ogg-1 hover:text-white hover:bg-indigo-950 text-2xl"
  );

  const [passwordRepeatText, setPasswordRepeatText] =
    useState("Repeat password");
  const [passwordRepeatStyles, setPasswordRepeatStyles] = useState(
    "text-white bg-teal-ogg-1 hover:text-white hover:bg-indigo-950 text-2xl"
  );

  const params = useParams();
  const refCode = params.id as string;
  var onSubmitClick = useCallback(() => {
    var registerInterface: IDoRegister = {
      password: "",
      email: "",
      securityText: "",
      referralCode: refCode,
      ecdhPrivate: "",
    };

    const passwd = passwordRef.current;
    const passwdRepeat = passwordRepeatRef.current;

    const passwordHash = DoPasswordOperations(
      passwd,
      passwdRepeat,
      setPasswordStyles,
      setPasswordRepeatStyles,
      setPasswordText,
      setPasswordRepeatText
    );

    passwordHash !== "" ? (registerInterface.password = passwordHash) : void 0; //password stuff ends here

    DoCheckMailValidity(emailRef.current, setMailText, setMailStyles) 
      ? (registerInterface.email = emailRef.current?.innerText as string)
      : void 0; //mail stuff ends here

    
    

  }, []);

  return (
    <div className="min-h-full ml-7 mr-7">
      <div className="pt-14 min-h-1/5">
        <Navbar collapse={true} />
      </div>
      <div className="w-full">
        <p className="text-5xl font-robotoSlab text-indigo-ogg flex justify-center pt-4">
          REGISTER
        </p>
      </div>

      <div className=" min-h-4/5 pt-25 ">
        <div className="flex flex-col">
          <div className="px-40  flex flex-row w-full space-x-[300px]">
            <div className="w-1/2">
              <div className="w-full">
                <ObeseBar
                  refPassed={emailRef}
                  height="min-h-[110px]"
                  color={mailStyles}
                  text={mailText}
                  contentEditable={true}
                />
              </div>
            </div>
            <div className="w-1/2">
              <div className="w-full">
                <ObeseBar
                  refPassed={passwordRef}
                  height="min-h-[110px]"
                  color={passwordStyles}
                  contentEditable={true}
                  text={passwordText}
                />
              </div>
            </div>
          </div>
          <div className="px-40  flex flex-row w-full space-x-[300px] ">
            <div className="w-1/2">
              <div className="w-full mt-6">
                <ObeseBar
                  refPassed={securityTextRef}
                  height="min-h-[370px]"
                  color="text-white bg-teal-ogg-1 hover:text-white hover:bg-indigo-950  text-2xl"
                  text="Enter arbitrary text not surpassing 32 characters, do save it somewhere secure and not lose it"
                  contentEditable={true}
                />
              </div>
            </div>
            <div className="w-1/2 mt-6 flex flex-col">
              <div>
                <ObeseBar
                  refPassed={passwordRepeatRef}
                  height="min-h-[110px]"
                  color={passwordRepeatStyles}
                  text={passwordRepeatText}
                  contentEditable={true}
                />
              </div>
              <div className="w-full mt-auto">
                <ObeseBar
                  refPassed={submitRef}
                  height="min-h-[110px]"
                  color="text-white bg-indigo-800 hover:text-white hover:bg-red-600 items-center justify-center text-3xl"
                  text="REGISTER"
                  onClick={onSubmitClick}
                  contentEditable={false}
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
