import { useSearchParams } from "react-router-dom";
import Navbar from "../../Navbar/Navbar";

export default function ErrorPage({ code }: { code?: number }) {
  let initializedCode = 404;
  const [searchParams, _] = useSearchParams();
  const paramCode = Number(searchParams.get("code"));
  const message = searchParams.get("message")

  if (paramCode !== null && !isNaN(paramCode)) {
    initializedCode = Number(paramCode);
  }

  initializedCode = code !== undefined ? code : 404

  return (
    <div className="w-full">
      <div className="pt-14 min-h-1/5">
        <Navbar collapse={true} />
      </div>
      <div className="w-full pt-15 ">
        <p className="text-7xl font-robotoSlab text-red-500 flex justify-center pt-4">
          ERROR {initializedCode} <br /> {message}
        </p>
      </div>
    </div>
  );
}
