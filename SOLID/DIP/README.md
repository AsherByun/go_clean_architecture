## DIP (의존성 역전)

### 개념
- 변동성이 큰 구체 클래스를 참조하지 말라 -> 대신 추상 인터페이스를 참조하라 (일반적으로 추상 팩토리)
- 변동성이 큰 구체 클래스로부터 파생하지 말라
- 구체 함수를 오버라이드 하지마라 -> 추상 메서드를 각각 구현하라
- 구체적이며 변동성이 크다면 절대로 그 이름을 언급하지 마라

### 스프링
- 개발자가 원래 직접 코드를 작성한 후 코드의 흐름을 제어했음 <-> 반대로 프레임워크는 직접 흐름을 제어함 (IOC)
- 스프링에서 IOC를 구현한 것이 DI  
- 인터페이스를 사용하여 의존성을 주입받아 의존성 역전을 구현하여 결합도를 낮추고, 유지보수를 쉽게만들어줌