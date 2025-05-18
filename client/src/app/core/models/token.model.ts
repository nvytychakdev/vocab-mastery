export class Token {
  constructor(
    public jwtToken: string,
    public expiresAt: string
  ) {}

  isExpired() {
    return new Date(this.expiresAt).valueOf() < Date.now();
  }

  toString(): string {
    return this.jwtToken;
  }
}
