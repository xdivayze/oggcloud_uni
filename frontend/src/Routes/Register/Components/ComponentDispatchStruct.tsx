import { Dispatch, SetStateAction, useRef, useState } from "react";

export default function ComponentDispatchStruct( 
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
