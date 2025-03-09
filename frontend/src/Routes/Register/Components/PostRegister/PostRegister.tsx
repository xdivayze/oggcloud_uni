import ErrorPage from "../../../ErrorPage/ErrorPage"
import UserCreated from "./UserCreated"


export default function PostRegister({code, success}:{code:number, success: boolean}) {
    if (!success) {
        
        return <ErrorPage code={code} />
    }
    return <UserCreated />
}
