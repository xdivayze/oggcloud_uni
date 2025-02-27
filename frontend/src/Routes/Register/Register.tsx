import { useCallback, useRef } from "react";
import Navbar from "../../Navbar/Navbar";
import ObeseBar from "./Components/ObeseBar";
import { useParams } from "react-router-dom";

export default function Register() {
  var emailRef = useRef("");
  var passwordRef = useRef("");
  var passwordRepeatRef = useRef("");
  var securityTextRef = useRef("");
  var submitRef = useRef("");

  const { refcode } = useParams()

  var onSubmitClick = useCallback(() => {
    console.log("hi man");

    var email = emailRef.current
    console.log(`${email != null ? email.innerText : "err"}`); //TODO continue from here adding fetch requests 

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
                  color="text-white bg-teal-ogg-1 hover:text-white hover:bg-indigo-950  text-2xl"
                  text="Enter your email(e.g. example@example.org)"
                  contentEditable={true}
                />
              </div>
            </div>
            <div className="w-1/2">
              <div className="w-full">
                <ObeseBar
                  refPassed={passwordRef}
                  height="min-h-[110px]"
                  color="text-white bg-teal-ogg-1 hover:text-white hover:bg-indigo-950  text-2xl"
                  contentEditable={true}
                  text="Enter a password not over 9 characters"
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
                  color="text-white bg-teal-ogg-1 hover:text-white hover:bg-indigo-950   text-2xl"
                  text="Repeat password"
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

async function checkRefCode(referral: string): Promise<boolean> { //TODO
  return true
}
