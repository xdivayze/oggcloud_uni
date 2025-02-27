import { useEffect, useState } from "react";
import ExpandedNavbar from "./Components/ExpandedNavbar";

export default function Navbar({ collapse }: { collapse: boolean }) {
  const [collapsed, setCollapsed] = useState(false);
  const [transitionStates, setTransitionStates] = useState("min-h-14");
  const [textStatement, setTextStatement] = useState("");
  //TODO FIX THIS NAVBAR!!
  useEffect(() => {
    handleCollapse();
  }, []);
  var handleCollapse = () => {
    if (collapse) {
      if (collapsed) {
        setTextStatement("text-2xl");
        setTransitionStates(
          "transition-all duration-500 ease-in-out  min-h-14 "
        );
        setCollapsed(false);
      } else {
        setTextStatement("text-[0px]");
        setTransitionStates(
          "transition-all duration-500 ease-in-out  min-h-1 max-h-4 "
        );
        setCollapsed(true);
      }
    }
  };

  var rendered = (
    <ExpandedNavbar
      height=""
      showText={textStatement}
      additional={`${transitionStates}`}
    />
  );

  return (
    <div
      className="w-13/14 flex ml-7 min-h-14 "
      onMouseEnter={handleCollapse}
      onMouseLeave={handleCollapse}
    >
      {rendered}
    </div>
  );
}
