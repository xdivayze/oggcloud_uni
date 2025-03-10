import { JSX, useCallback, useEffect, useRef, useState } from "react";
import ObeseBar from "./ObeseBar";
import Navbar from "../../../Navbar/Navbar";
import { ComponentDispatchStruct, DoRegister } from "../Services/Register";
import { useParams } from "react-router-dom";
import { DoPasswordOperations } from "../Services/PasswordServices";
import { DoCheckMailValidity } from "../Services/MailServices";
import GenerateKeys from "../Services/KeyGenerationService";
import { IDoRegister, StatusCodes } from "../Services/utils";

export default function RegisterSuccess() {
  const emailRef = useRef<HTMLDivElement>(null);
  const passwordRef = useRef<HTMLDivElement>(null);
  const passwordRepeatRef = useRef<HTMLDivElement>(null);
  const securityTextRef = useRef<HTMLDivElement>(null);
  const submitRef = useRef<HTMLDivElement>(null);

  const [passwordText, setPasswordText] = useState(
    "Enter a password not over 9 characters"
  );
  const [passwordStyles, setPasswordStyles] = useState(
    "text-white bg-teal-ogg-1 hover:text-white hover:bg-indigo-950 text-2xl"
  );

  const passwordCompStruct: ComponentDispatchStruct = {
    setStyle: setPasswordStyles,
    setText: setPasswordText,
    compRef: passwordRef,
    originalStyle: useRef(passwordStyles).current,
  };

  const [mailText, setMailText] = useState(
    "Enter your email(e.g. example@example.org)"
  );
  const [mailStyles, setMailStyles] = useState(
    "text-white bg-teal-ogg-1 hover:text-white hover:bg-indigo-950 text-2xl"
  );

  const mailCompStruct: ComponentDispatchStruct = {
    setStyle: setMailStyles,
    setText: setMailText,
    compRef: emailRef,
    originalStyle: useRef(mailStyles).current,
  };

  const [passwordRepeatText, setPasswordRepeatText] =
    useState("Repeat password");
  const [passwordRepeatStyles, setPasswordRepeatStyles] = useState(
    "text-white bg-teal-ogg-1 hover:text-white hover:bg-indigo-950 text-2xl"
  );

  const passwordRepeatCompStruct: ComponentDispatchStruct = {
    setStyle: setPasswordRepeatStyles,
    setText: setPasswordRepeatText,
    compRef: passwordRepeatRef,
    originalStyle: useRef(passwordRepeatStyles).current,
  };

  const [securityTextText, setSecurityTextText] = useState(
    "Enter arbitrary text not surpassing 32 characters, do save it somewhere secure and not lose it"
  );
  const [securityTextStyles, setSecurityTextStyles] = useState(
    "text-white bg-teal-ogg-1 hover:text-white hover:bg-indigo-950  text-2xl"
  );

  const securityTextCompStruct: ComponentDispatchStruct = {
    setStyle: setSecurityTextStyles,
    setText: setSecurityTextText,
    compRef: securityTextRef,
    originalStyle: useRef(securityTextStyles).current,
  };

  const params = useParams();
  const refCode = params.id as string;
  const [render, setRender] = useState<JSX.Element>(<></>);

  const onSubmitClick = useCallback(() => {
    const registerInterface: IDoRegister = {
      password: "",
      email: "",
      referralCode: refCode,
      ecdhPublic: "",
      secText: ""
    };

    const passwordHash = DoPasswordOperations(
      passwordCompStruct,
      passwordRepeatCompStruct
    );

    passwordHash !== "" ? (registerInterface.password = passwordHash) : void 0; //password stuff ends here

    DoCheckMailValidity(mailCompStruct)
      ? (registerInterface.email = emailRef.current?.innerText as string)
      : void 0; //mail stuff ends here

    GenerateKeys(securityTextCompStruct)
      .then(({ code, ecdhPub }) => {
        code === StatusCodes.Success
          ? (() => {
            registerInterface.secText = securityTextRef.current?.innerText as string;
              registerInterface.ecdhPublic = ecdhPub as string; 
              DoRegister(registerInterface, setRender);
            })()
          : void 0; //encryption stuff ends here
      })
      .catch((e: Error) => {
        console.error(e);
      });
  }, []);

  //set the default render and update once register request is sent
  const defaultRender = //this is probably not a good idea...
    (
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
                    color={securityTextStyles}
                    text={securityTextText}
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
  useEffect(() => {
    setRender(defaultRender);
  }, []);

  return render;
}
