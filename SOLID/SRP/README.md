# SRP (Single Responsibility Principle)

### 개념
- 원래 생각 -> 하나의 메소드, 객체는 하나의 행위만 수행해야 한다
- 클린아키텍쳐 책 -> 오직 하나의 액터에 대해서만 책임저야 한. (ex employee data - cto, coo, cfo...)


### 우발적 중복
- 서로다른 주체가 하나의 클래스를 공유하여 사용할 때 한 주체에 의해 변경시 다른 주체에 영향을 줄 수 있다
- ex) 단순 메소드 변경, 테이블 스키마 조정

### 병합