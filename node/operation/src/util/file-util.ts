import fs from 'fs';

type PutText = string | number;

export const stat = async (path: string): Promise<void> => {
    return new Promise(resolve => {
        fs.stat(path, (err, data) => {
            if (err) throw err;
            resolve();
        });
    });
};

export const read = async (path: string): Promise<string> => {
    return new Promise(resolve => {
        fs.readFile(path, 'utf-8', (err, data) => {
            if (err) throw err;
            resolve(data);
        });
    });
};

export const write = async (path: string, text: PutText): Promise<void> => {
    return new Promise(resolve => {
        fs.writeFile(path, String(text), (err) => {
            if (err) throw err;
            resolve();
        });
    });
};

export const removeFile = async (path: string): Promise<void> => {
    return new Promise(resolve => {
        fs.unlink(path, (err) => {
            if (err) throw err;
            resolve();
        });
    });
};

export const makedir = async (path: string): Promise<void> => {
    return new Promise(resolve => {
        fs.mkdir(path, { recursive: true }, (err) => {
            if (err) throw err;
            resolve();
        });
    });
};

export const removeDir = async (path: string): Promise<void> => {
    return new Promise(resolve => {
        fs.rmdir(path, { recursive: true }, (err) => {
            if (err) throw err;
            resolve();
        })
    })
};
