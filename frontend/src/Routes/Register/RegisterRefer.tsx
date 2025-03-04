import { useRef } from "react";
import ObeseBar from "./Components/ObeseBar";
import Navbar from "../../Navbar/Navbar";
import { useNavigate, useSearchParams } from "react-router-dom";

export default function RegisterRefer({
  submitText,
  submitColor,
}: {
  submitText: string;
  submitColor: string;
}) {
  const referCodeRef = useRef<HTMLDivElement>(null);
  const submitRef = useRef<HTMLDivElement>(null);
  const navigate = useNavigate();

  const [searchParams, _] = useSearchParams()
    const successCode = searchParams.get("code")
    if ( successCode === "-1") {
      submitText = "403 FORBIDDEN"
      submitColor = "bg-red-700"
      
    }

  const onSubmitClick = () => {
    const refCode: string = referCodeRef.current === null ? "" : referCodeRef.current.innerText;
    if (refCode == "") {
      console.log("error, ref code null");
      return;
    }
    const refCodeTrimmed = refCode.trim();
    if (/\s/.test(refCodeTrimmed) || refCodeTrimmed.length != 64) {
      console.log("error, ref code includes whitespace or is not 64 characters");
      return;
    }
    navigate(`/register/${refCode.trim()}`);
  };
  return (
    <div className="w-full">
      <div className="pt-14 min-h-1/5">
        <Navbar collapse={true} />
      </div>
      <div className="w-full">
        <p className="text-5xl font-robotoSlab text-indigo-ogg flex justify-center pt-4">
          REGISTER
        </p>
      </div>
      <div className="w-full flex justify-center min-h-screen pt-25 ">
        <div className="flex flex-col w-1/2">
          <div className="w-full">
            <ObeseBar
              refPassed={referCodeRef}
              height="min-h-[110px]"
              color="text-white bg-teal-ogg-1 hover:text-white hover:bg-indigo-950  text-2xl"
              text="Enter your referral code"
              contentEditable={true}
            />
          </div>
          <div className="w-full mt-6">
            <ObeseBar
              refPassed={submitRef}
              height="min-h-[110px]"
              color={`${submitColor} text-white hover:text-white hover:bg-red-600 items-center justify-center text-3xl`}
              text={submitText}
              onClick={onSubmitClick}
              contentEditable={false}
            />
          </div>
        </div>
      </div>
    </div>
  );
}
