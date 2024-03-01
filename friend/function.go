package friend

import "errors"

type Friend struct {
	id int8
	nama string
	alamat string
	pekerjaan string
	alasan string
}

type FriendsCollection struct {
	friends []Friend
}

func NewFriend(id int8, nama, alamat, pekerjaan, alasan string) Friend {
	return Friend{id: id, nama: nama, alamat: alamat, pekerjaan: pekerjaan, alasan: alasan}
}

func (fc *FriendsCollection) AddFriend(data Friend) {
	fc.friends = append(fc.friends, data)
}

func (fc *FriendsCollection) GetAllFriend() ([]Friend) {
	return fc.friends
}

func (fc *FriendsCollection) GetFriend(id int8) (Friend, error) {
	var result Friend

	for _, f := range fc.friends {
		if f.id == id {
			result = f
			return result, nil
		}
	}

	return result, errors.New("Friend not found")
}