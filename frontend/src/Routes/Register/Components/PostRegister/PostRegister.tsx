import Err from "./Err"
import UserCreated from "./UserCreated"

export default function PostRegister({success}:{success:boolean}) {
    if (!success) {
        return <Err />
    }
    return <UserCreated />
}
