import { hkdf } from "@noble/hashes/hkdf";

import { createHash } from "crypto";
import { sha256 } from "@noble/hashes/sha256";

import { Buffer } from "buffer/";
import { ERR_MODE_STYLES, StatusCodes } from "./utils";
import { ComponentDispatchStructType } from "../Components/ComponentDispatchStruct";

export const ECDH_PRIVATE_STORAGE_FIELD = "ecdhPrivate";
export const AES_MASTERKEY_STORAGE_FIELD = "masterKey";
export const SEED_FIELD = "seed";

export default async function GenerateKeys(
  securityTextCompStruct: ComponentDispatchStructType
): Promise<{
  code: StatusCodes;
  masterKey?: Buffer | null;
  ecdhPub?: string | null;
}> {
  const securityTextE = securityTextCompStruct.getRef();

  securityTextCompStruct.setStyles(securityTextCompStruct.originalStyles);

  if (securityTextE.current === null) {
    securityTextCompStruct.setText(StatusCodes.ErrNull);
    securityTextCompStruct.setStyles(ERR_MODE_STYLES);
    return { code: StatusCodes.ErrNull };
  }

  const securityText = securityTextE.current.innerText;
  if (securityText.trim().length < 16) {
    securityTextCompStruct.setText(StatusCodes.ErrSecurityTextTooShort);
    securityTextCompStruct.setStyles(ERR_MODE_STYLES);
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
    SEED_FIELD,
    securityTextCompStruct.getRefContent().innerText
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
