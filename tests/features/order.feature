Feature: change order status
    Scenario: change order to received
        When I change from PENDENTE to RECEBIDO
        Then the result should equal true

    Scenario: change order to in prepare
        When I change from RECEBIDO to EM_PREPARO
        Then the result should equal true

    Scenario: change order to done
        When I change from EM_PREPARO to FINALIZADO
        Then the result should equal true

    Scenario: change order to finished
        When I change from FINALIZADO to ENTREGUE
        Then the result should equal true

    Scenario: change order to invalid status
        When I change from FINALIZADO to RECEBIDO
        Then the result should equal false
