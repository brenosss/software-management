import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:frontend/entities/languageInfo.dart';

Future<List<LanguageInfo>> fetchLanguageInfo() async {
  final response = await http.get(Uri.http("localhost:8080", "/languages"));

  if (response.statusCode == 200) {
    // If the server did return a 200 OK response,
    // then parse the JSON.
    final Map<String, dynamic> responseBody = jsonDecode(response.body);
    Iterable list = responseBody['data'];
    return list.map((item) => LanguageInfo.fromJson(item)).toList();
  } else {
    // If the server did not return a 200 OK response,
    // then throw an exception.
    throw Exception('Failed to load languages info');
  }
}
