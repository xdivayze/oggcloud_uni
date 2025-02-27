import { Link } from "react-router-dom";

export default function TextParts() {
  return (
    <div className="w-full text-5xl font-robotoSlab text-white flex justify-center flex-col leading-15">
      <p>Cloud Storage,</p>
      <p>
        <span className="text-orange-ogg">Private</span> and
        <span className="text-yellow-ogg"> Exclusive.</span>
      </p>
      <div className="flex flex-col pt-1 ">
        <p className="">We believe that by</p>
        <p>putting a user cap, we</p>
        <p>
          can <span className="text-blue-ogg-0">improve</span> our user
        </p>
        <p>experience.</p>
      </div>

      <div className="pt-1 flex flex-col max-w-fit">
        <p className="">
          <span className="text-indigo-ogg underline hover:scale-110 hover:text-blue-700 transition-all duration-200">
            {" "}
            <Link to={"/register"}>Join now</Link>{" "}
          </span>
          with your
        </p>
        <p className="flex justify-center pr-3">referral</p>
      </div>
    </div>
  );
}
