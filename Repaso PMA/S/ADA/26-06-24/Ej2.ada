Procedure Ej2 is

Task type Worker;

arrWorkers: array (1..10) of Worker;

Task Coordinador is 
    entry recibirSuma(suma: IN integer);
End Coordinador;

Task Coordinador is
    sumaT: real;
Begin
    loop i in 1..10 do
        accept recibirSuma(suma: IN integer) do
            sumaT := sumaT + suma;
        end recibirSuma;
        sumaT := sumaT /100000;
        promedioTotal := promedioTotal + sumaT;
    end loop;

End Coordinador;

Task Body Worker is
    arrValores: array (1..100000) of integer;
    sumaLocal: integer;
begin
    for i in 1..100000 loop
        sumaLocal := sumaLocal + arrValores[i];
    end loop
    Coordinador.recibirSuma(sumaLocal);
end worker;

begin 

end Ej2;