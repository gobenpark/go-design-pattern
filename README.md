# GO Design Patterns


## Creational Patterns

1. [SingleTon](https://github.com/BumwooPark/go-design-pattern/tree/master/singleton)
: 글로벌하게 하나의 오브젝트에 대한 접근이 필요할 경우

2. [Builder](https://github.com/BumwooPark/go-design-pattern/tree/master/builder)
: 빌더 패턴은 일반적인 구성 알고리즘을 사용하여 예측이 불가능한 object의 수를 유지하는데 도움이 됨 
인터페이스의 변경이 있다면 모든 빌더에 영향을 미치고 메소드가 추가 될경우 어색할수 있기에 해당 알고리즘이 안정될 것을 확신하지 못한다면 이패턴을 피하자

3. [Factory Method](https://github.com/BumwooPark/go-design-pattern/tree/master/factory_method)

4. [Abstract Factory](https://github.com/BumwooPark/go-design-pattern/tree/master/abstract_factory)
: factory method는 `Factory Method`보다 좀더 큰 범주에 속함  `Factory Method` 가 객체를 반환한다면 
`Abstract Factory`는 `factory` 자체를 반환하고 그 `factory`를 통해 객체를 받아서 처리를 한다.

5. [Object Pool](https://github.com/BumwooPark/go-design-pattern/tree/master/object_pool)
6. [ProtoType Pattern](https://github.com/BumwooPark/go-design-pattern/tree/master/prototype)


## Structural Patterns

1. [Composite Pattern](https://github.com/BumwooPark/go-design-pattern/tree/master/composite)
: 트리구조의 객체 계층구조를 만듬  상속과 다중상속의 문제를 해결함 즉 열차를 타고 수영하는 수영선수를 만들경우

2. [Adapter Pattern](https://github.com/BumwooPark/go-design-pattern/tree/master/adapter)
: 공개/폐쇄 원칙을 유지 하며 코드가 새로운 기능에는 개방 되어있고 수정에는 폐쇄되어있는 방식 레거시 코드를 새로운코드로 연결 하는 방식 
대표적으로 Go 에서는 http.HandlerFunc , ServeHTTP 등을 찾아보고 HandlerFunc의 경우 함수로 interface를 구현하는 신선한 방식
두개의 호환되지 않는 

3. [Bridge Pattern](https://github.com/BumwooPark/go-design-pattern/tree/master/bridge)
: 브릿지 패턴의 목적은 자주 변경되는 구조체에 유연성을 가져오는 것 , 메서드의 입력과 출력을 알면 코드를 많이 모르는 상태에서 코드를 쉽게 수정이 가능하다

4. [Proxy Pattern](https://github.com/BumwooPark/go-design-pattern/tree/master/proxy)
: 사용자에게 권한을 부여하거나 데이터베이스에 대한 액세스 권한을 부여하는 것과 같은 중간 작업이 필요한 유형을 중심으로 프록시를 래핑하는것이 좋음 (캐시 구현)

5. [Decorator Pattern](https://github.com/BumwooPark/go-design-pattern/tree/master/decorator)
: 런타임시 추가적인 기능을 붙일때 주로 쓰이는 부분이 서버의 미들웨어 부분

6. [Flyweight Pattern](https://github.com/BumwooPark/go-design-pattern/tree/master/flyweight)
: 데이터를 공유하여 메모리를 절약하는 패턴으로 공통으로 사용되는 객체는 한 번만 생성되고 공유를 통해 풀(Pool)에 의해 관리, 사용된다.


## Behavioral Patterns

1. [Strategy Pattern](https://github.com/BumwooPark/go-design-pattern/tree/master/strategy)
: 특정 기능을 구현하기 위한 몇 가지 알고리즘 제공 , 모든 유형이 동일한 기능을 다른 방식으로 실행하지만 클라이언트는 영향을 받지 않는다.
factory 패턴과 go의 struct 임베딩 + 인터페이스 덕타이핑  
a 구조체에 b 구조체를 임베딩 할 경우 a 에 속해있는 것처럼 b필드에 엑세스가 가능하다 이부분이 상속과 비슷 하지만 차이점이 있는데 
b에 해당하는 함수가 필요할경우 a를 전달할 수 없다 하지만 b가 구현하는 인터페이스를 받아들이는 함수가 있다면 a를 전달 할 수 있다.
![image](https://github.com/BumwooPark/go-design-pattern/blob/master/strategy/strategy.jpg?raw=true)

2. [Chain Responsibility Pattern](https://github.com/BumwooPark/go-design-pattern/tree/master/chain_responsibility)
: 각 채인은 단일책임 원칙을 준수한다, 해당 패턴은 로깅체인을 구현하는것으로 프로그램의 출력을 둘 이상의 io.Writer인터페이스에 쓰는 것이다.
로깅을 할때 콘솔/파일/remote 서버에 남기는 기능이 있을때 3번의 함수콜을 하는것보다 체이닝을 통해서 한번에 처리하는게 더 유용하다
behavior와 state의 처리를 런타임에 하는 것으로 FSM(Finite State Machine)을 만드는데 널리 사용이 된다.
주의할점은 체이닝중 중간에 실패할경우에 대한 처리가 필요 