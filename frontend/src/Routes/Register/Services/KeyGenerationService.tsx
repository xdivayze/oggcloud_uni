import { hkdf } from "@noble/hashes/hkdf";
import {
  ComponentDispatchStruct,
} from "./Register";
import { createHash } from "crypto";
import { sha256 } from "@noble/hashes/sha256";

import { Buffer } from "buffer/";
import { ERR_MODE_STYLES, StatusCodes } from "./utils";

const ECDH_PRIVATE_STORAGE_FIELD = "ecdhPrivate";
const AES_PRIVATE_STORAGE_FIELD = "masterKey";

export default async function GenerateKeys(
  securityTextCompStruct: ComponentDispatchStruct
): Promise<{
  code: StatusCodes;
  masterKey?: Buffer | null;
  ecdhPub?: string | null;
}> {
  const {
    compRef: securityTextE,
    setStyle: setSecurityTextStyles,
    setText: setSecurityTextText,
    originalStyle,
  } = securityTextCompStruct;
  setSecurityTextStyles(originalStyle);

  if (securityTextE.current === null) {
    setSecurityTextText(StatusCodes.ErrNull);
    setSecurityTextStyles(ERR_MODE_STYLES);
    return { code: StatusCodes.ErrNull };
  }

  const securityText = securityTextE.current.innerText;
  if (securityText.trim().length < 16) {
    setSecurityTextText(StatusCodes.ErrNull);
    setSecurityTextStyles(ERR_MODE_STYLES);
    return { code: StatusCodes.ErrSecurityTextTooShort };
  }
  const securityHash = createHash("sha256")
    .update(securityText.trim())
    .digest();

  const derivedKey = hkdf(sha256, securityHash, "", "ECDH Key", 32);

  const ecdhSecret = await crypto.subtle.generateKey(
    { name: "ECDH", namedCurve: "P-256" },
    true,
    ["deriveKey", "deriveBits"]
  );

  const exportedECDHPublic = await crypto.subtle.exportKey(
    "spki",
    ecdhSecret.publicKey
  );
  let ECDHPublicB64Pem = "";

  ECDHPublicB64Pem = formatAsPem(
    Buffer.from(exportedECDHPublic).toString("base64")
  );

  const exportedECDHPrivate = await crypto.subtle.exportKey(
    "pkcs8",
    ecdhSecret.privateKey
  );

  window.localStorage.setItem(
    AES_PRIVATE_STORAGE_FIELD,
    Buffer.from(derivedKey).toString("hex")
  );
  window.localStorage.setItem(
    ECDH_PRIVATE_STORAGE_FIELD,
    Buffer.from(exportedECDHPrivate).toString("hex")
  );

  return {
    code: StatusCodes.Success,
    masterKey: Buffer.from(derivedKey),
    ecdhPub: ECDHPublicB64Pem,
  }; //return deterministic symmetric key for session key encryption
}

function formatAsPem(s: string): string {
  let finalString = "-----BEGIN PUBLIC KEY-----\n";
  while (s.length > 0) {
    finalString += s.substring(0, 64) + "\n";
    s = s.substring(64);
  }
  finalString = finalString + "-----END PUBLIC KEY-----";
  return finalString;
}
