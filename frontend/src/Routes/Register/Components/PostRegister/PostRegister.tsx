import { useSearchParams } from "react-router-dom";

export default function PostRegister(  ) {
    const [searchParams, _] = useSearchParams()
    let returnElement = <div></div>

    if (searchParams.get("code") as unknown as number !== 201) {
        
    }

    return returnElement
}
