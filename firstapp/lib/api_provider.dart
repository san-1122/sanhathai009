import 'package:http/http.dart' as http;
import 'dart:io';
import 'dart:async';
import 'dart:convert';

class ApiProvider {
  ApiProvider();

  String endPoint = 'http://192.168.90.202:8080';

  Future<http.Response> doLogin(String username, String password) async {
    var url = '$endPoint/login';
    var body = {
      "username": username,
      "password": password,
    };

    var loginResponse = await http.post(
      Uri.parse(url),
      headers: {
        HttpHeaders.contentTypeHeader: 'application/json',
      },
      body: jsonEncode(body),
    );

    return loginResponse;
  }

  Future<http.Response> doRegister(
      String email, String password, String fullname, String avatar) async {
    var url_r = '$endPoint/register';

    var body = {
      "username": email,
      "password": password,
      "fullname": fullname,
      "avatar": avatar,
    };

    var registerResponse = await http.post(
      Uri.parse(url_r),
      headers: {
        HttpHeaders.contentTypeHeader: 'application/json',
      },
      body: jsonEncode(body),
    );

    return registerResponse;
  }
}
