typedef long int i32;

void caesarEncrypt(i32 *plaintext, i32 plaintextLength, i32 key) {
  for (int i = 0; i < plaintextLength; i++) {
    plaintext[i] = (plaintext[i] + key) % 26;
  }
}


void caesarDecrypt(i32 *cyphertext, i32 cyphertextLength, i32 key) {
  for (int i = 0; i < cyphertextLength; i++) {
    cyphertext[i] = (cyphertext[i] -  key) % 26;
  }
}
