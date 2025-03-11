import { Dispatch, RefObject, SetStateAction, useRef, useState } from "react";

export default class ComponentDispatchStruct {
  private  setStyleState: Dispatch<SetStateAction<string>>;
  public  styles: string;
  private  setTextState: Dispatch<SetStateAction<string>>;
  public  text: string;
  public  originalStyles: string;
  private  compRef: RefObject<HTMLDivElement | null>;
  public constructor(initialStyles: string, initialText: string) {
    [this.styles, this.setStyleState] = useState(initialStyles);
    [this.text, this.setTextState] = useState(initialText);
    this.originalStyles = useRef(initialStyles).current;
    this.compRef = useRef<HTMLDivElement>(null);
  }
  public setStyles(newStyles: string) {
    this.setStyleState(newStyles);
  }
  public setText(newText: string) {
    this.setTextState(newText);
  }
  public getRef() {
    return this.compRef;
  }
  public getRefContent() {
    if (this.compRef === null) {
        console.error("problem occurred")
      throw new Error("reference null");
    } else 
    {return this.compRef.current as unknown as HTMLDivElement};
  }
  
}
