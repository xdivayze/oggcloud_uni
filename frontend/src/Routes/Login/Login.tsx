import { useEffect,  } from "react";
import Navbar from "../../Navbar/Navbar";
import ObeseBar from "../Register/Components/ObeseBar";

import ComponentDispatchStruct from "../Register/Components/ComponentDispatchStruct";

export default function Login() {
  useEffect(() => {}, []); //TODO check for saved sign-in

  const defaultStyles =
    "text-white bg-teal-ogg-1 hover:text-white hover:bg-indigo-950 text-2xl";

  const emailCompStruct = new ComponentDispatchStruct(
    defaultStyles,
    "Enter your email(e.g. example@example.org)"
  );

  // const [emailText, setEmailText] = useState(
  //   "Enter your email(e.g. example@example.org)"
  // );
  // const [emailStyles, setEmailStyles] = useState(
  //   "text-white bg-teal-ogg-1 hover:text-white hover:bg-indigo-950 text-2xl"
  // );
  // const emailRef = useRef<HTMLDivElement>(null);

  // const emailCompStruct: ComponentDispatchStruct = {
  //   setStyle: setEmailStyles,
  //   setText: setEmailText,
  //   compRef: emailRef,
  //   originalStyle: useRef(emailStyles).current,
  // };

  const passwordCompStruct = new ComponentDispatchStruct(
    defaultStyles,
    "Enter your password"
  );

  // const [passwordText, setPasswordText] = useState("Enter your password");
  // const [passwordStyles, setPasswordStyles] = useState(
  //   "text-white bg-teal-ogg-1 hover:text-white hover:bg-indigo-950 text-2xl"
  // );
  // const passwordRef = useRef<HTMLDivElement>(null);
  // const passwordCompStruct: ComponentDispatchStruct = {
  //   setStyle: setPasswordStyles,
  //   setText: setPasswordText,
  //   compRef: passwordRef,
  //   originalStyle: useRef(passwordStyles).current,
  // };

  // const secTextCompStruct: ComponentDispatchStruct = {

  // }

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
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
