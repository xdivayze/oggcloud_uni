import { Dispatch, SetStateAction, useRef } from "react";
import ObeseBar from "../../Register/Components/ObeseBar";

export default function SubmitButton({
  setSave,
  onSubmitClick,
}: {
  setSave: Dispatch<SetStateAction<boolean>>;
  onSubmitClick: () => void;
}) {
  return (
    <>
      <ObeseBar
        refPassed={useRef(null)}
        height="min-h-[110px]"
        color="text-white bg-indigo-800 hover:text-white hover:bg-red-600 items-center justify-center text-3xl"
        text="REGISTER"
        onClick={onSubmitClick}
        contentEditable={false}
      />
      <div className="w-full pt-2 text-white text-xl font-robotoSlab">
        <label className="flex items-center space-x-2 pl-2">
          <input
            onChange={(e) => {
              setSave(e.target.checked);
            }}
            className="w-5 h-5 rounded focus-ring-blue-400"
            type="checkbox"
          />
          <span>Save user info</span>
        </label>
      </div>
    </>
  );
}
