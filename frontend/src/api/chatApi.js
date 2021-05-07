import axios from 'axios';

const postHeaders = {
	'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8',
};

export function searchChatRecords(searchData) {
	return axios.post('/api/v1/search/chatmessages', searchData, postHeaders);
}
