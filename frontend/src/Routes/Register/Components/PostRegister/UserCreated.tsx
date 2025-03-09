import { useRef } from "react";
import Navbar from "../../../../Navbar/Navbar";
import ObeseBar from "../ObeseBar";

export default function UserCreated({
  securityText,
}: {
  securityText: string;
}) {
  const seedRef = useRef<HTMLDivElement>(null);
  return (
    <div className="w-full">
      <div className="w-full pt-14 min-h-1/5">
        <Navbar collapse={true} />
      </div>
      <div className="w-full pt-28 ">
        <div className="w-full px-50 flex justify-center ">
          <ObeseBar
            color="bg-indigo-800 text-white hover:text-white hover:bg-red-600 items-center justify-center text-3xl"
            text={securityText}
            refPassed={seedRef}
            height="min-h-[370px]"
            contentEditable={false}
          />
        </div>
      </div>
    </div>
  );
}
