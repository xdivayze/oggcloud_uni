import { useEffect, useState } from "react";
import Navbar from "../../Navbar/Navbar";
import FetchPreviewIms, {
  FETCH_IMAGE_COUNT,
  FileData,
} from "./Service/FetchPreviewIms";
import Preview from "./Components/Preview";

export default function Profile() {
  const [previews, setImagePreviews] = useState<FileData[]>([]);
  useEffect(() => {
    for (let i = 0; i <= FETCH_IMAGE_COUNT; i++) {
      setImagePreviews((prevItems) => [...prevItems, FetchPreviewIms(i)]);
    }
  }, []);
  return (
    <div className="w-full px-7 flex flex-col h-screen py-14">
      <div className="min-h-1/7 w-full">
        <Navbar collapse={true} />
      </div>
      <div className="bg-teal-ogg-1 w-full flex-grow rounded-2xl">
        {previews.map((v, i) => {
          const pr = Preview(v);
          return <div key={i}>{pr.element}</div>;
        })}
      </div>
    </div>
  );
}
