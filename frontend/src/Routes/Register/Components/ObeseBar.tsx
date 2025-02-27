import React from "react";

export default function ObeseBar({
  height,
  color,
  text,
  refPassed,
  onClick,
  contentEditable,
}: {
  height: string;
  color: string;
  text: string;
  refPassed: React.RefObject<any>;
  onClick?: () => void;
  contentEditable: boolean
}) {
  return (
    <div
    onKeyDown={(e) => {
      if (e.key === "Enter") {
        console.log("e")
        e.preventDefault()
      }
    }}
      onClick={onClick}
      suppressContentEditableWarning={true}
      ref={refPassed}
      contentEditable={contentEditable}
      className={`w-full rounded-[30px] font-robotoSlab flex  transition-all duration-300 ${!contentEditable ? "cursor-default" : "pr-5 pl-5 pt-5"}   ${height} ${color}`}
    >
      {text}
    </div>
  );
}
