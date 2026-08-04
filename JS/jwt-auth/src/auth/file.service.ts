import { Injectable } from '@nestjs/common';
import path from 'node:path';
import { SignupDto } from './dto/signup.dto';
import { readFile, writeFile } from 'node:fs/promises';

type StoredUser = SignupDto & { refreshToken?: string };

@Injectable()
export class FileService {
  private readonly filePath = path.join(process.cwd(), 'users.json');

  async appendFile(user: SignupDto) {
    let users: StoredUser[] = [];

    try {
      const file = await readFile(this.filePath, 'utf-8');
      users = JSON.parse(file) as StoredUser[];
    } catch {
      // File doesn't exist yet
    }

    users.push(user);

    await writeFile(this.filePath, JSON.stringify(users, null, 2), 'utf-8');

    return { username: user.username };
  }

  async readFile(): Promise<StoredUser[]> {
    try {
      const file = await readFile(this.filePath, 'utf-8');
      const users = JSON.parse(file) as StoredUser[];

      return users;
    } catch {
      return [];
    }
  }

  async updateUser(username: string, updates: Partial<StoredUser>) {
    const users = await this.readFile();
    const updated = (users ?? []).map((u) =>
      u.username === username ? { ...u, ...updates } : u,
    );

    await writeFile(this.filePath, JSON.stringify(updated, null, 2));
    return updated.find((u) => u.username === username);
  }
}
