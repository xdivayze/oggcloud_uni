import { useEffect } from "react";
import Navbar from "../../Navbar/Navbar";
import ObeseBar from "../Register/Components/ObeseBar";

import ComponentDispatchStruct from "../Register/Components/ComponentDispatchStruct";
import { ObeseBarDefaultStyles } from "../Register/Services/utils";

export default function Login() {
  useEffect(() => {}, []); //TODO check for saved sign-in

  const defaultStyles = ObeseBarDefaultStyles;

  const emailCompStruct = new ComponentDispatchStruct(
    defaultStyles,
    "Enter your email(e.g. example@example.org)"
  );

  const passwordCompStruct = new ComponentDispatchStruct(
    defaultStyles,
    "Enter your password"
  );

  const securityTextCompStruct = new ComponentDispatchStruct(ObeseBarDefaultStyles,
    "Enter your security txt"
    
  );

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
            <div className="w-1/2 mt-6">
              <div className="w-full mt-auto"></div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
