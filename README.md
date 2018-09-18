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

7. [Composite Pattern](https://github.com/BumwooPark/go-design-pattern/tree/master/composite)
: 트리구조의 객체 계층구조를 만듬  상속과 다중상속의 문제를 해결함 즉 열차를 타고 수영하는 수영선수를 만들경우 