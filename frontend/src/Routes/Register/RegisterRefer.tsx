import { useRef } from "react";
import ObeseBar from "./Components/ObeseBar";
import Navbar from "../../Navbar/Navbar";
import { useNavigate } from "react-router-dom";

export default function RegisterRefer() {
    const refercodeRef = useRef("");
    const submitRef = useRef("")
    const navigate = useNavigate()
    const onSubmitClick = () => {
        const refcode: string = refercodeRef.current.innerText
        if (refcode == "") {
            console.log("error, refcode null")
            return
        }
        const refcodeTrimmed = refcode.trim() 
        if (/\s/.test(refcodeTrimmed)) {
            console.log("error, refcode includes whitespace")
            return
        }
        navigate(refcode.trim())
    }
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
                            refPassed={refercodeRef}
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
                            color="text-white bg-indigo-800 hover:text-white hover:bg-red-600 items-center justify-center text-3xl"
                            text="REGISTER"
                            onClick={onSubmitClick}
                            contentEditable={false}
                        />
                    </div>
                </div>
            </div>
        </div>
    );
}
