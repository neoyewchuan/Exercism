import 'package:hello_world/hello_world.dart';
import 'package:test/test.dart';

void main() {
  test('Hi, world!', () {
    expect(HelloWorld().hello(), equals('Hello, World!'));
  });
}
