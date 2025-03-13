import { JSX, useCallback, useEffect, useRef, useState } from "react";
import ObeseBar from "./ObeseBar";
import Navbar from "../../../Navbar/Navbar";
import { DoRegister } from "../Services/Register";
import { useParams } from "react-router-dom";
import { DoPasswordOperations } from "../Services/PasswordServices";
import { DoCheckMailValidity } from "../Services/MailServices";
import GenerateKeys from "../Services/KeyGenerationService";
import {
  IDoRegister,
  ObeseBarDefaultStyles,
  StatusCodes,
} from "../Services/utils";
import ComponentDispatchStruct from "./ComponentDispatchStruct";
import PostRegister from "./PostRegister/PostRegister";

export default function RegisterSuccess() {
  const submitRef = useRef<HTMLDivElement | null>(null);

  const passwordCompStruct = ComponentDispatchStruct(
    ObeseBarDefaultStyles,
    "Enter a password not over 9 characters"
  );

  const mailCompStruct = ComponentDispatchStruct(
    ObeseBarDefaultStyles,
    "Enter your email(e.g. example@example.org)"
  );
  const passwordRepeatCompStruct = ComponentDispatchStruct(
    ObeseBarDefaultStyles,
    "Repeat password"
  );
  const securityTextCompStruct = ComponentDispatchStruct(
    ObeseBarDefaultStyles,
    "Enter arbitrary text not surpassing 32 characters, do save it somewhere secure and not lose it"
  );

  const params = useParams();
  const refCode = params.id as string;
  const [render, setRender] = useState<JSX.Element>(<></>);

  const [submitted, setSubmitted] = useState(false);
  const [responseStatus, setResponseStatus] = useState(0);

  const onSubmitClick = useCallback(() => {
    const registerInterface: IDoRegister = {
      password: "",
      email: "",
      referralCode: refCode,
      ecdhPublic: "",
      secText: "",
    };

    const passwordHash = DoPasswordOperations(
      passwordCompStruct,
      passwordRepeatCompStruct
    );

    if (passwordHash === "") {
      console.error("err")
      return;
    }
    registerInterface.password = passwordHash; //password stuff ends here

    if (DoCheckMailValidity(mailCompStruct)) {
      
      registerInterface.email = mailCompStruct.getRefContent().innerHTML;
    } else return ;//mail stuff ends here


    GenerateKeys(securityTextCompStruct)
      .then(({ code, ecdhPub }) => {
        if (code === StatusCodes.Success) {
          registerInterface.secText =
            securityTextCompStruct.getRefContent().innerText;
          registerInterface.ecdhPublic = ecdhPub as string;
          DoRegister(registerInterface)
            .then((v) => {
              setResponseStatus(v as unknown as number);
            })
            .catch((e) => console.error(e));
          setSubmitted(true);
        } //encryption stuff ends here
      })
      .catch((e: Error) => {
        console.error(e);
        throw e;
      });
  }, []);
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
                    refPassed={mailCompStruct.getRef()}
                    height="min-h-[110px]"
                    color={mailCompStruct.styles}
                    text={mailCompStruct.text}
                    contentEditable={true}
                  />
                </div>
              </div>
              <div className="w-1/2">
                <div className="w-full">
                  <ObeseBar
                    refPassed={passwordCompStruct.getRef()}
                    height="min-h-[110px]"
                    color={passwordCompStruct.styles}
                    contentEditable={true}
                    text={passwordCompStruct.text}
                  />
                </div>
              </div>
            </div>
            <div className="px-40  flex flex-row w-full space-x-[300px] ">
              <div className="w-1/2">
                <div className="w-full mt-6">
                  <ObeseBar
                    refPassed={securityTextCompStruct.getRef()}
                    height="min-h-[370px]"
                    color={securityTextCompStruct.styles}
                    text={securityTextCompStruct.text}
                    contentEditable={true}
                  />
                </div>
              </div>
              <div className="w-1/2 mt-6 flex flex-col">
                <div>
                  <ObeseBar
                    refPassed={passwordRepeatCompStruct.getRef()}
                    height="min-h-[110px]"
                    color={passwordRepeatCompStruct.styles}
                    text={passwordRepeatCompStruct.text}
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
  }, [
    securityTextCompStruct.text,
    passwordCompStruct.text,
    passwordRepeatCompStruct.text,
    mailCompStruct.text,
  ]); //TODO side effect doesn't trigger when same value is assigned 
  return !submitted ? (
    render
  ) : (
    <PostRegister
      success={responseStatus === 201 ? true : false}
      code={responseStatus}
    />
  );
}
