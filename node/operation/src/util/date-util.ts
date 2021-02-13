export const now = (): Date => new Date(Date.now());

export const toHHMM  = (date: Date):string  => {
    return (
        ('00' + String(date.getHours())).slice(-2) +
        ('00' + String(date.getMinutes())).slice(-2)
    );
};

export const toYYYYMMDD = (date: Date): string => {
    return (
        String(date.getFullYear()) +
        String(('0' + String(date.getMonth())).slice(-2)) +
        String(('0' + String(date.getDate())).slice(-2))
    );
};
