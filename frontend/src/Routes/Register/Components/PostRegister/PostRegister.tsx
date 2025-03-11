import ErrorPage from "../../../ErrorPage/ErrorPage";
import UserCreated from "./UserCreated";

import { SEED_FIELD } from "../../Services/KeyGenerationService";

export default function PostRegister({
  code,
  success,
}: {
  code: number;
  success: boolean;
}) {
  const secText = window.localStorage.getItem(SEED_FIELD) as unknown as string;
  if (!success) {
    return <ErrorPage code={code} />;
  } else if (secText !== undefined) {
    return <UserCreated securityText={secText} />;
  } else {
    return <ErrorPage code={500} />;
  }
}
