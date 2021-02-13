import { DateUtil } from "../";

describe('DateUtil', () => {
    const date = new Date(Date.now());
    it('toHHMM', () => {
        expect(DateUtil.toHHMM(date).length).toBe(4);
    });
    
    it('toYYYYMMDD', () => {
        expect(DateUtil.toYYYYMMDD(date).length).toBe(8);
    });
});