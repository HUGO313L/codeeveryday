$(function() {
	setTimeout(function() {
		console.log("123");
	    console.log($('#my').attr("href"));
	    $('#my').attr('href', 'http://www.huawei.com');
	}, 0);// ���ó�0�����ӡ���Ϊhuawei������Ϊ1�����ӡ���Ϊbaidu�����˵����phantomjs����Ƚű���ִ����ɲ�ִ��open�Ļص���
});