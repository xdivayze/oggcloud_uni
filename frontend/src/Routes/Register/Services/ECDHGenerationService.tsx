import { hkdf } from "@noble/hashes/hkdf";
import { StatusCodes } from "./Register";
import { createHash, randomBytes } from "crypto";
import { sha256 } from "@noble/hashes/sha256";
import { p256, secp256r1 } from "@noble/curves/p256";

export default async function GenerateECDH(securityText: string): Promise<{
  code: StatusCodes;
  data: Buffer<ArrayBuffer> | null;
}> {
  if (securityText.trim().length < 16) {
    return { code: StatusCodes.ErrSecurityTextTooShort, data: null };
  }
  const salt = randomBytes(32);
  const securityHash = createHash("sha256")
    .update(securityText.trim())
    .digest();
  const derivedKey = hkdf(sha256, securityHash, salt, "ECDH Key", 32);
  const privateKey = p256.utils.isValidPrivateKey(derivedKey)
    ? derivedKey
    : (() => {
        const n = secp256r1.CURVE.n;
        const keyBigInt = BigInt(
          "0x" + Buffer.from(derivedKey).toString("hex")
        );

        const validKeyBigInt = keyBigInt % n;

        const validKey = new Uint8Array(32);
        for (let i = 0; i < 32; i++) {
          validKey[31 - i] = Number((validKeyBigInt >> BigInt(8 * i)) & 0xffn);
        }
        return validKey;
      })();
  return { code: StatusCodes.Success, data: Buffer.from(privateKey) };
}
