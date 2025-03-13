import { Dispatch, SetStateAction, useRef, useState } from "react";

export default function ComponentDispatchStruct( //TODO doesn't update state ui test stuff in another project and implement maybe a reducer or switch to redux
  initialStyles: string,
  initialText: string
): ComponentDispatchStructType {
  const [styles, setStyles] = useState(initialStyles);
  const [text, setText] = useState(initialText);
  const originalStyles = useRef(initialStyles).current;
  const compRef = useRef<HTMLDivElement>(null);

  return {
    styles,
    text,
    setStyles,
    setText,
    originalStyles,
    getRef: () => compRef,
    getRefContent: () => {
      if (!compRef.current) throw new Error("reference null");
      return compRef.current;
    },
  };
}

export interface ComponentDispatchStructType {
  styles: string;
  text: string;
  setStyles: Dispatch<SetStateAction<string>>;
  setText: Dispatch<SetStateAction<string>>;
  originalStyles: string;
  getRef: () => React.RefObject<HTMLDivElement | null>;
  getRefContent: () => HTMLDivElement;
};
