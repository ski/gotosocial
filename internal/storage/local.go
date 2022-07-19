/*
   GoToSocial
   Copyright (C) 2021-2022 GoToSocial Authors admin@gotosocial.org

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package storage

import (
	"context"
	"io"
	"net/url"

	"codeberg.org/gruf/go-store/kv"
)

type Local struct {
	KVStore *kv.KVStore
}

func (l *Local) Get(ctx context.Context, key string) ([]byte, error) {
	return l.KVStore.Get(key)
}

func (l *Local) GetStream(ctx context.Context, key string) (io.ReadCloser, error) {
	return l.KVStore.GetStream(key)
}

func (l *Local) PutStream(ctx context.Context, key string, r io.Reader) error {
	return l.KVStore.PutStream(key, r)
}

func (l *Local) Put(ctx context.Context, key string, value []byte) error {
	return l.KVStore.Put(key, value)
}

func (l *Local) Delete(ctx context.Context, key string) error {
	return l.KVStore.Delete(key)
}

func (l *Local) URL(ctx context.Context, key string) *url.URL {
	return nil
}
