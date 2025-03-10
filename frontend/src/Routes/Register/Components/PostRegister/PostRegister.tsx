import ErrorPage from "../../../ErrorPage/ErrorPage";
import UserCreated from "./UserCreated";

export default function PostRegister({
  code,
  success,
  secText,
}: {
  code: number;
  success: boolean;
  secText?: string;
}) {
  
  if (!success) {
    return <ErrorPage code={code} />;
  } else if (secText !== undefined) {
    return <UserCreated securityText={secText} />;
  } else {
    return <ErrorPage code={500} />;
  }
}
