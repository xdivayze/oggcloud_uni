import { hkdf } from "@noble/hashes/hkdf";
import { sha256 } from "@noble/hashes/sha256";
import { createHash, randomBytes } from "crypto";
import { Buffer } from "buffer/";
import { AUTH_CODE_FIELDNAME } from "../Routes/Login/Service/constants";
import { MAIL_FIELDNAME } from "../Routes/Register/Services/MailServices";

export interface CredentialPool {
  password: string;
  masterKey: string;
  email: string;
  ecdhPrivate: string;
}

export const SESSION_ID_FIELDNAME = "cookieSessionID";

export default async function SaveLogin(
  credentials: CredentialPool,
  authCode: string
) {
  //TODO add endpoint for cookie stuff in backend
  const { data, iv } = await encryptCookieData(credentials).catch(
    (e: Error) => {
      console.error(e);
      throw e;
    }
  );
  const endpoint = "/api/user/remember-me";
  const req = await fetch(endpoint, {
    method: "POST",
    headers: {
      [AUTH_CODE_FIELDNAME]: authCode,
      [MAIL_FIELDNAME]: credentials.email,
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      data: data,
      iv: iv,
    }),
  })
    .catch((e: Error) => {
      throw e;
    })
    .then((v) => v.json())
    .then((v) => {
      if (v.status !== 201) {
        throw new Error("status non-201");
      }
      if (!(SESSION_ID_FIELDNAME in v)) {
        throw new Error(
          SESSION_ID_FIELDNAME + " doesn't exist on response body"
        );
      }
      return v[SESSION_ID_FIELDNAME];
    }); //TODO save the cookie

  return;
}

async function encryptCookieData(credentials: CredentialPool) {
  const passwordHash = createHash("sha256")
    .update(credentials.password)
    .digest("hex");
  const kek = await hkdf(sha256, passwordHash, "", "KEK", 32);
  const importedKek = await window.crypto.subtle.importKey(
    "raw",
    kek,
    { name: "AES-GCM" },
    false,
    ["encrypt", "decrypt"]
  );
  const iv = randomBytes(16);
  const encryptedData = await window.crypto.subtle
    .encrypt(
      { name: "AES-GCM", iv },
      importedKek,
      new TextEncoder().encode(
        JSON.stringify({
          ecdhPrivate: credentials.ecdhPrivate,
          masterKey: credentials.masterKey,
        })
      )
    )
    .catch((e: Error) => {
      throw e;
    });
  const ivString = iv.toString("hex");
  const encryptedDataString = Buffer.from(encryptedData).toString("hex");
  return { data: encryptedDataString, iv: ivString };
}
