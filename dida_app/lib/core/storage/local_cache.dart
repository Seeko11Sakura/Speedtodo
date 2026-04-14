import 'package:shared_preferences/shared_preferences.dart';

class LocalCache {
  const LocalCache(this._preferences);

  final SharedPreferences _preferences;

  Future<void> saveBaseUrl(String value) async {
    await _preferences.setString('api_base_url', value);
  }

  String? readBaseUrl() {
    return _preferences.getString('api_base_url');
  }
}
