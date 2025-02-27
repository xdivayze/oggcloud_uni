import Navbar from "../../Navbar/Navbar";
import SpotifyPart from "./Components/SpotifyPart";
import TextParts from "./Components/TextPart";

export default function Home() {
  return (
    <div className="min-h-full ml-7 mr-7">
      <div className="pt-14 min-h-1/5">
        <Navbar collapse={true} />
      </div>

      <div className=" min-h-4/5 pt-14 ">
        <div className="px-35 w-full flex flex-row ">
          <div className="w-1/2 ">
            <TextParts />
          </div>
          <div className="w-1/2 ">
            <SpotifyPart />
          </div>
        </div>
      </div>
    </div>
  );
}
