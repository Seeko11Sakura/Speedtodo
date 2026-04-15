import 'package:dio/dio.dart';

final api = Dio(BaseOptions(
  baseUrl: 'http://10.0.2.2:8080',
  connectTimeout: const Duration(seconds: 5),
  receiveTimeout: const Duration(seconds: 5),
));
