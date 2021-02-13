import { FileUtil } from '../'

describe('FileUtil', () => {
    it('stat', async () => {
        expect(await FileUtil.stat('fixtures/read.txt')).resolves;
    });
    it('read', async () => {
        expect(await FileUtil.read('fixtures/read.txt')).toBe('This is a file.');
    });
    it('write', async () => {
        expect(await FileUtil.write('fixtures/write.txt', 'This is a text')).resolves;
    });
    it('removeFile', async () => {
        expect(await FileUtil.removeFile('fixtures/write.txt')).resolves;
    });
    it('makedir', async () => {
        expect(await FileUtil.makedir('fixtures/write.txt')).resolves;
    });
    it('removeDir', async () => {
        expect(await FileUtil.removeDir('fixtures/write.txt')).resolves;
    });
});