export function toRFC3339(date: string): string {
    return new Date(date).toISOString();
}
