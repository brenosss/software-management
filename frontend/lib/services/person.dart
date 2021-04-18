import 'dart:convert';
import 'package:frontend/entities/person.dart';
import 'package:http/http.dart' as http;

Future<List<Person>> fetchPersonList() async {
  final response = await http.get(Uri.http("localhost:8080", "/person"));

  if (response.statusCode == 200) {
    final Map<String, dynamic> responseBody = jsonDecode(response.body);
    Iterable list = responseBody['data'];
    return list.map((item) => Person.fromJson(item)).toList();
  } else {
    throw Exception('Failed to load person list');
  }
}

Future<Person> fetchPersonDetail(String personId) async {
  final response =
      await http.get(Uri.http("localhost:8080", "/person/$personId"));

  if (response.statusCode == 200) {
    final Map<String, dynamic> responseBody = jsonDecode(response.body);
    Map item = responseBody['data'];
    return Person.fromJson(item);
  } else {
    print(response.body);
    throw Exception('Failed to load person detail');
  }
}
