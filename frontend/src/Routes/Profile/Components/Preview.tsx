
import { FileData } from "../Service/FetchPreviewIms";
import FetchOriginalImage from "../Service/FetchOriginalImage";
import { JSX } from "react";

export function PreviewElem() {
  return <div></div>;
}

export default function Preview(data: FileData): {
  element: JSX.Element;
  fetchOriginal: () => FileData;
} {
  return {
    element: <PreviewElem />,
    fetchOriginal: () => {
      return FetchOriginalImage(data.id);
    },
  };
}
